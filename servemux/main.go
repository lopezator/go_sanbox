package main

import (
	"net/http"
	"log"
)

func main() {
	// Crear ServeMux vacío
	mux := http.NewServeMux()

	//Crear un handler (en este caso, uno de la stdlib de Golang, RedirectHandler)
	//Lo que hace, es redirigir todo lo que le llega a la URL example.org
	rh := http.RedirectHandler("http://example.org", 307)

	//"Atacheamos" ese handler al ServeMux que hemos creado arriba
	mux.Handle("/foo", rh)

	log.Println("Listening...")

	//Lo dejamos escuchando en el puerto 3000
	http.ListenAndServe(":3000", mux)
}
