## Notes on Go Server

## Go Server

* We start by creating a mux instance. http.NewServeMux(). Mux is a request multiplexer. It matches the URL of each
  incoming request against a list of registered patterns and calls the handler for the pattern that most closely matches
  the URL.
* Mux is also nothing but a simple handler, which delegates control to secondary registered handlers. It is like
  chaining requests.
* Every handler must implement the ServerHTTP method.
* Go provides various default handlers like RedirectHandler(url string, code int).
* Each mux.Handle requires a route path and a handler for this path.

### Typical Implementation

```go
//internal go
type Handler interface {
    ServeHTTP(w ResponseWriter, r *http.Request)
}

func main() {
    mux := http.NewServerMux()
    th := TimeHandler{format: time.RFC1123}
    mux.Handle("/time", th)
    http.ListenAndServer(":3000", mux)
}

type TimeHandler struct {
    format string
}

// ServeHTTP any type which implements the Handler interface is a Handler. The type can be passed in place of the interface.
func (th TimeHandler) ServeHTTP(w ResponseWriter, r *http.Request) {
    tn := time.Now().Format(th.format)
    w.Write([]byte("time is: " + tn)
}
```

### Can we avoid the redundant struct?
- Creating a type just for the handler looks redundant, can we avoid this?
- In below code we statically define the format.
- We require to pass on a handler to mux, here timeHandler is not a handler. We can solve this by coercing it into a Handler by using http.HandlerFunc
- The second approach removes the boiler plate of coercing as well, mux provides a nice wraper around it using mux.HandleFunc

```go
func main() {
    mux := http.NewServerMux()
    th := http.HandlerFunc(timeHandler)
    // we need to pass on a handler to mux
    mux.Handle("/time", th)
    http.ListenAndServer(":3000", mux)

}

func timeHandler(w ResponseWriter, r *http.Request) {
    tn := time.Now().Format(time.RFC1123)
    w.Write([]byte("time is: " + tn)
}
```

```go
func main() {
    mux := http.NewServerMux()
    mux.HandlerFunc("/time", timeHandler)
    http.ListenAndServer(":3000", mux)
}

func timeHandler(w ResponseWriter, r *http.Request) {
    tn := time.Now().Format(time.RFC1123)
    w.Write([]byte("time is: " + tn)
}
```

### How to pass variables to handlers?

- Closures, we can pass variables to a func and make it return a handler. The function closes in on the variable and hence called the closure.
- In the below approach the return is implicitly converted to http.HandlerFunc type. We can also perform return http.HandlerFunc(fn) or http.HandlerFunc(func(ResponseWriter, *http.Request))

```go
func main() {
    //.,...
    mux.Handle("/time", timeHandler)
    //...
}

func timeHandler(format string) http.HandlerFunc {
    // closes in the variable format. 
    return func(w ResponseWriter, r *http.Request) {
        tn := time.Now().Format(format)
        w.Write([]byte("time is: " + tn)
    }
}
```



References

1. Handler & ServerMutexes: https://www.alexedwards.net/blog/an-introduction-to-handlers-and-servemuxes-in-go
2. Handling Json in Request Body: https://www.alexedwards.net/blog/how-to-properly-parse-a-json-request-body 
