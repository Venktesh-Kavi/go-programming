## File Operations

Explore on the various options to read and write to a file in go.

### Opening/Creating a file

```go
    f, err := os.Open(fileName) // returns a file (a Reader)

    // os.Open()/os.Create() is an abstraction for os.OpenFile() were we open with mode as O_RDONLY for read or O_CREATE/O_RDWR/O_TRUNC and perm as 0 and 0666 respectively.
```

### Reading a file

Reading from a file offers various options. We can use

```go
    buf := []byte{}
    f.Read(buf) // returns the no. of bytes read and any error encountered. if full file is read returns error as io.EOF

    bufio.NewReader() // returns a reader. initialises a reader for io.Reader object with default buffer size of 4096 bytes (4KB)

    bufio.NewScanner() // returns a scanner which has options for a split function and the max token size. Default max token size 64 * 1024 (64KB).

    bytes, err := ioUtil.ReadFile("foo.yaml") // reads an entire file in one go. Use with caution as everything is in memory.

    fileList, err := ioUtil.ReadDir("~/personal") // reads an entire directory
```

### Which to use when

* Using file.Read() can be done for many use cases. 
* bufio.Scanner() provides convenience in reading files, one can provide a custom split function. Eg.., bufio.ScanLines() inbuilt split function advances the pointer to the next line for every iterations and the scanner provides method to extract the byte array at every iteration between the start and the end pointers.
* Use bufio.Scanner() when we know the buffer size statically.
An example can be parsing a csv file. More or less the row size is fixed,though there can be a single row with all values filled in. 
* Use bufio.Reader() when you require fine grained control over error handling, max buffer token size or run sequential scans use the bufio.Reader().
* Use utility functions like ioUtil.ReadFile to read files in one go for quick use cases.


## References

- https://chriswilcox.dev/blog/2024/04/09/Scan-vs-Read-in-bufio.html
- https://kgrz.io/reading-files-in-go-an-overview.html
- https://yourbasic.org/golang/io-reader-interface-explained/#:~:text=Basics-,The%20io.,read%20a%20stream%20of%20bytes.&text=Read%20reads%20up%20to%20len,error%20when%20the%20stream%20ends.
- https://www.reddit.com/r/golang/comments/7opf2y/reading_file_concurrently_using_bufioscanner_need/
