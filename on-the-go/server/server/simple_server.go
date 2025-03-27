package server

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// StartServer Refer server_notes.md
func StartServer() {
	log.Println("Starting server")
	mux := http.NewServeMux()
	mux.HandleFunc("/time", timeHandler(time.RFC1123))
	mux.Handle("/foo", http.RedirectHandler("https://www.google.com", http.StatusMovedPermanently))
	http.ListenAndServe(":8080", mux)
}

func timeHandler(format string) http.HandlerFunc {
	// closure
	return func(w http.ResponseWriter, r *http.Request) {
		tn := time.Now().Format(format)
		w.Write([]byte("The time is: " + tn))
	}
}

func personHandler(w http.ResponseWriter, r *http.Request) {
	p := &Person{}
	buf := make([]byte, 1024)
	r.Body.Read(buf)
	err := json.Unmarshal(buf, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
