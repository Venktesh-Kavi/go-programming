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
	http.Handle("/foo", DecorateLog(http.HandlerFunc(HandleGetFoo(HandlerStruct{"GET"}))))
	log.Printf("starting server in port: %s", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Printf("unable to start server in port: %s", port)
	} else {
		log.Printf("started http server in port: %s", port)
	}
}

func DecorateLog(fn http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Started processing request:", r.Method, r.URL.Path)
		defer log.Println("Completed processing request:", r.Method, r.URL.Path)
		fn.ServeHTTP(w, r)
	})
}

func wrappedWriter(w io.Writer, errStr string) {
	_, e := fmt.Fprintf(w, errStr)
	if e != nil {
		fmt.Printf("unable to write to http response writer buffer, err: %v", e)
	}
}
func filter(verb string) bool {
	mm := map[string]struct{}{"POST": {}, "HEAD": {}, "PUT": {}, "OPTIONS": {}}
	if _, ok := mm[verb]; !ok {
		return true
	}
	return false
}
