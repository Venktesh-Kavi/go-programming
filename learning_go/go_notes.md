## General Learning on Go

* Building Go Programs in any directory
* assume the go file is present in test/helloworld/hello.go
  ``` go
  // build the whole package
  go build -o hello test/helloworld/
  
  // if a singular file with main is present
  go build -o hello test/helloworld/hello.go
  ```
* build nested submodules which have a go.mod file
  ``` go
  go build ./...
  ```