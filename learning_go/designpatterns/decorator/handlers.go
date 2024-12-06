package decorator

import (
	"net/http"
)

type HandlerStruct struct {
	supportedMethod string
}

// HandlerGetFoo
func HandleGetFoo(h HandlerStruct) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if !filter(h.supportedMethod) {
			w.WriteHeader(500)
			wrappedWriter(w, r.Method+" http method not supported")
		} else {
			wrappedWriter(w, "Got FOO")
		}
	}
}

func HandlePostFoo(w http.ResponseWriter, r *http.Request) {

}

func HandleFallBack(handler http.Handler) {
}
