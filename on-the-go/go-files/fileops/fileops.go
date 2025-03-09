package fileops

import (
	"fmt"
	"io"
	"log"
	"os"
)

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
