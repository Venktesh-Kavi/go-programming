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

## Handling JSON Payload in Requests

### Basic Implementation
- The below approach is fine for proto-type use cases. If its for production use, few things can be improved:

1. Not all errors returned by Decode() are as a result of bad request from a client. Specifically if it is json.InvalidUnMarshallError error, it is because of an incorrect type sent to the Decode() method. It should result in Internal server error rather than bad request
2. The error messages returned by Decode() method aren't ideal for the clients. Some give too much detail about the underlying program eg.., (cannot marshal json to person.name of type string). Others are un-descriptive like unexpected io.EOF.
3. A client can include extra unexpected fields in their JSON, these fields are silently ignored without the client being notified about it or erroring. We can fix this by using decoder's DisallowUnknownFields() method.
4. There is no upper limit on the size of the request payload which can be read from the server. Limiting this would help the server from exhausting resources, if the client maliciously sends a large request. This can be limited using http.MaxBytesRead().
5. We are not verifying the content type. (application/json)
6. The decoder being used is designed decode streams of json like {"foo": 123}{"bar": 432} or {"foo": 123}{}. But as per the above code only the first object in the request needs to be parsed. If a client sends multiple objects we need to notify them.

For point 6, there are two ways to address this either call the Decoder() method again for the second time & make sure it returns an io.EOF. Or avoid using Decode() altogether and use json.Unmarshall()m it would return an error if multiple objects are present. But it doesn't solve point (3) does stop from sending unknown fields.


Refer improved_handler.go code for a better handler.


```go

// LoanDetailsReq request payload
type LoanDetailsReq struct {
    clientId                    string  `json: "clientId"`
    clientApplicationId         string  `json: "clientAppId"
    partnershipApplicationId    string 
}

func main() {
    mux := http.NewServerMux()
    mux.HandleFunc("/person", personHandler)
    http.ListenAndServe(":3000", mux)
}

func personHandler(w ResponseWriter, r *http.Request) {
    var p Person
    err := json.NewDecoder(r.Body).Decode(&p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Fprintf formats and prints it to the provided writer.
    fmt.Fprintf(w, "Person: %v", p)

}

```

### Go Structs and JSON 

#### Encoding
- Encodes a struct to a json encoded stream and writes them to the provided the stream supplied.
- To encode the struct fields should be exported. (Why: Go uses pacakge as its primary mechanism of encapsulation. Since the json encoder is defined in a different package, if not exported it will be unable to see the fields). Passing unexported fields will result in empty json object without any error.
- The output stream can be a writable stream, here we are writing to stdout. In webservers we will use http response stream to write the encoded json.
- If the Encode() method is not able to convert it into a json, it throws an error. Incase maps are used to convert to json, the map key must be a string.
- Encoder() can handle all kind of data types (maps, slices, pointers, structs).
- Make sure not to have circular references eg.., a struct pointing to itself. The go garbage collector can handle two structs pointing to itself. But the encoder() can get caught in infinite cycle and the program can hang up.
- 

```go
p := Person{Name: "Venktesh", Id: "12312"}
enc := json.NewEncoder()
if err := enc.Encode(p); err != nil {
    fmt.Printf("error in encoding %v\n", err)
}
```

Customising Encoding using Tags and Options

```go

type Person struct {
    Name    string  `json: "name,omitempty"`
    Id      string  `json: "id,omitempty"`
}

func main() {
    p := Person{Name: "Venktesh", Id: "23123"}
    enc := json.NewEncoder()
    if err := enc.Encode(p); err != nil {
        fmt.Printf("error in encoding %v\n", p)
    }
}
```

#### Marshalling
- Encoding, encodes the type to a json and puts it to an writable stream.
- Sometimes we might need the encoded json in-memory to perform some modifications or pass it along to some non-stream oriented function (eg.., can be a database call).

```go
type Rectagle struct {
    Top     int
    Left    int
    Width   int
    Height  int
}

r := Rectagle{10, 20, 30, 40}
buf, err := json.Marshall(r)
```
- json.Marshall() returns a slice of bytes and potentially an error. The byte slice can be manipulated however we like or passed along.

#### Decoding
- Converting a JSON stream to a struct.

```go
r := &Rectangle{}

dec := json.NewDecoder(os.Stdin)
if err := dec.Decode(&r); err != nil {
    fmt.Printf("error in decoding %v\n", err)
}
```
- The above code reads a json string provided from the os's stdin (sys shell).
- The decoder will match the json keys to struct fields and will silently ignore ones which it cannot find.

#### Unmarshalling
- Similar to marshalling which converts a struct to json buffer byte slice rather than a stream.
- Unmarshalling converts a json byte slice to a struct.

```go
r := &Rectangle{}
s := `{"top": 10, "left": 20, "width": 30, "height": 40}`
buf := []byte(s)
if err := json.Unmarshall(buf, r); err != nil {
    fmt.Printf("error in unmarshalling %v\n", err)
}
```



## References

1. Handler & ServerMutexes: https://www.alexedwards.net/blog/an-introduction-to-handlers-and-servemuxes-in-go
2. Handling Json in Request Body: https://www.alexedwards.net/blog/how-to-properly-parse-a-json-request-body 
3. https://drstearns.github.io/tutorials/gojson/#:~:text=Go%20maps%20that%20use%20a,and%20pointers%20to%20other%20structs.
