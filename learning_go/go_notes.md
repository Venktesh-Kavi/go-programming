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

### Go Http Package

* Interfaces with these single methods are very powerful in go, examples are io.Reader, io.Writer,
  http.Handler.
* Any type which implements the ServeHttp method follows the Handler interface.
* Typically, when we want to do this we might create a struct and create a method to comply with the
  interface.
* Here we are using function type which represents the same contract as the ServeHttp method.
* So the http.HandlerFunc(pattern, fn func(ResponseWriter, *Request)) can essential be used instead
  of using http.Handler(pattern, Handler)

``` go
type Handler interface {
  ServeHttp(ResponseWriter, *Request)
}

type HandlerFunc func(ResponseWriter, *Request)

func (h HandlerFunc) ServeHttp(w ResponseWriter, r *Request) {
}
```