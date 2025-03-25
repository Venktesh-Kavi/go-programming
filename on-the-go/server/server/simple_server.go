package server

import (
	"log"
	"net/http"
	"time"
)

func StartServer() {
	log.Println("Starting server")
	mux := http.NewServeMux()
	mux.HandleFunc("/time", timeHandler)
	mux.Handle("/foo", http.RedirectHandler("https://www.google.com", http.StatusMovedPermanently))
	http.ListenAndServe(":8080", mux)
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	tn := time.Now().Format(time.RFC1123)
	w.Write([]byte("The time is: " + tn))
}
