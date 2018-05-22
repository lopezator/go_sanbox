package main

import (
	"net/http"
	"log"
	"time"
)

func timeHandler(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(time.RFC1123)
	w.Write([]byte("The time is: " + tm))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/time", timeHandler)

	log.Println("Listening...")
	http.ListenAndServe(":3000", mux)
}

