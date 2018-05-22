package main

import (
	"net/http"
	"time"
	"log"
)

type timeHandler struct {
	format string
}

// Para construir un custom handler
// Simplemente necesitamos un tipo que cumpla la interface handler
// Que se hace implementando el método ServeHTTP con este signature.
// +info: https://golang.org/src/net/http/server.go?s=2610:2673#L72
func (th *timeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(th.format)
	w.Write([]byte("The time is: " + tm))
}

func main() {
	mux := http.NewServeMux()

	th := &timeHandler{format: time.RFC1123}
	mux.Handle("/time", th)

	log.Println("Listening...")
	http.ListenAndServe(":3000", mux)
}

