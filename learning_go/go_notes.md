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

### New Keyword

* the `new` keyword in go allocates memory for a variable of the specified type and returns the
  pointer to the memory location.

``` go
func main() {
    t := new(Test)
    fmt.Println(t)
}
type Test struct {
    name    string
    age     int
}
```

* In the above example the o/p: `&{ 0}`. Indicates that it is a pointer

### Usage of Yield in Go

* Yield statement allows to return the execution context to the caller.
* Typically used for function that return an iterable context in go
* References:
    * [Ref](https://bbengfort.github.io/2016/12/yielding-functions-for-iteration-golang/) 