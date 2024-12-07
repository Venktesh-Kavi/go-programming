package decorator

import (
	"net/http"
)

// HandleGetFoo handles /foo route GET method
func HandleGetFoo(w http.ResponseWriter, r *http.Request) {
	wrappedWriter(w, "hello from foo route")
}
