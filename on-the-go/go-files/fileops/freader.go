package fileops

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"runtime"
	"sync"
	"time"
)

const DEFAULT_BUF_SIZE = 100

type chunk struct {
	size   int
	offset int64
}

/*
ReadFile simplicitic way to read file. Reads the file into the given buffer.
*/
func ReadFile(fn string) {
	fmt.Printf("reading file: %s\n", fn)
	f, err := os.Open(fn)
	if err != nil {
		log.Fatalf("unable to open file: %v\n", err)
	}

	finfo, _ := f.Stat()
	fmt.Printf("file stats: %v\n", finfo)

	buf := make([]byte, finfo.Size())
	fmt.Printf("created a buffer with size: %d\n", finfo.Size())

	for {
		_, err := f.Read(buf)
		if err == io.EOF {
			fmt.Println("reached end of file, it is an error though!")
			break
		}
		fmt.Println(string(buf))
	}
}

func ConvinientReader(fn string) {
	fmt.Printf("reading file: %s, using buffered scanner\n", fn)
	f := OpenFile(fn)
	sc := bufio.NewScanner(f)
	sc.Split(bufio.ScanLines) // Split accepts a splitFunc. Any func which splitFunc can go here.
	for sc.Scan() {
		fmt.Println(sc.Text())
	}
}

func FineGrainedReader(fn string) {
	fmt.Printf("reading file: %s, using low level abstraction reader\n", fn)
	f := OpenFile(fn)
	r := bufio.NewReader(f)

	for {
		b, err := r.ReadBytes(' ')
		if err != nil {
			if err != io.EOF {
				log.Fatalf("error in reading: %v\n", err)
			} else {
				log.Println("Succssfuly read file, reached EOF")
			}
			break
		}
		fmt.Println(string(b))
	}
}

/*
ReadFileConcurrently chunks the file and spawns go routines for each chunk.
*/
func ReadFileConcurrently(fn string) {
	numCPU := runtime.NumCPU()
	fmt.Printf("Reading file concurrently, File: %s, CPU: %d\n", fn, numCPU)
	f := OpenFile(fn)

	fs := int(getFileStats(f))

	chunks, concurrency := createChunks(fs)
	readChunks(f, chunks, concurrency)
}

func readChunks(f *os.File, chunks []chunk, concurrency int) {
	var wg sync.WaitGroup
	wg.Add(concurrency)
	for idx := range concurrency {
		// span go routines
		s := chunks[idx].size
		o := chunks[idx].offset
		buf := make([]byte, s)
		go func() {
			defer wg.Done()
			fmt.Println("Gopher >>> ", idx)
			time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
			_, err := f.ReadAt(buf, o) // drain the buffer
			if err != nil {
				if err != io.EOF {
					fmt.Printf("error reading chunk: %d\n", idx)
				} else if err == io.EOF {
					fmt.Println("successfuly completed reading file")
				}
			}
			fmt.Printf("Gopher: %d, Read: %s\n\n", idx, string(buf))
		}()
	}
	wg.Wait() // wait till all go go routines are done
}

func createChunks(fs int) ([]chunk, int) {
	concurrency := fs / DEFAULT_BUF_SIZE
	chunks := make([]chunk, DEFAULT_BUF_SIZE)

	for i := range concurrency {
		chunks[i] = chunk{DEFAULT_BUF_SIZE, int64(i * DEFAULT_BUF_SIZE)}
	}

	if r := concurrency % DEFAULT_BUF_SIZE; r != 0 {
		chunks = append(chunks, chunk{r, int64(concurrency * DEFAULT_BUF_SIZE)})
		concurrency++
	}

	fmt.Println("file size: , determined concurrency: ", fs, concurrency)
	return chunks, concurrency
}

func OpenFile(fn string) *os.File {
	f, err := os.Open(fn)
	if err != nil {
		log.Fatalf("unable to open file: %v", err)
	}
	return f
}

func getFileStats(f *os.File) int64 {
	info, err := f.Stat()
	if err != nil {
		log.Fatalf("unable to read stats from file: %v", err)
	}
	return info.Size()
}
