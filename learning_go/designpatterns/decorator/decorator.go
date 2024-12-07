package decorator

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

/* Typically decorator pattern is used when an interface/function is exposed by a third party library.
We want to decorate it with addition functionality to be used by downstream functions or activities.
*/

// http.HandlerFunc follow the decorator pattern. Any function which take an input fn x and returns x by doing decorations internally is a decorator.
const port = ":9096"

func InitDecoratorServer(wg *sync.WaitGroup) {
	defer wg.Done()
	// func(ResponseWriter, *Request)
	http.HandleFunc("/foo", Decorate(HandleGetFoo, "GET"))
	log.Printf("starting server in port: %s", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Printf("unable to start server in port: %s", port)
	} else {
		log.Printf("started http server in port: %s", port)
	}
}

func Decorate(fn http.HandlerFunc, method string) func(http.ResponseWriter, *http.Request) {
	return DLog(DFilter(fn, method))
}

func DFilter(fn func(w http.ResponseWriter, r *http.Request), method string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			// cannot use Fatalf results in server getting stopped. How to return
			w.WriteHeader(http.StatusBadRequest)
			wrappedWriter(w, fmt.Sprintf("request method not supported, received: %s, expected: %s\n", r.Method, method))
			return
		}
		fn(w, r)
	}
}
func DLog(fn func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("started processing request, method: %s, path: %s", r.Method, r.URL.Path)
		defer log.Printf("completed processing request, path: %s", r.URL.Path)
		fn(w, r)
	}
}

func wrappedWriter(w io.Writer, errStr string) {
	_, e := fmt.Fprintf(w, errStr)
	if e != nil {
		fmt.Printf("unable to write to http response writer buffer, err: %v", e)
	}
}
