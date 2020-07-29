package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/saludar", saludar)
	//subir un servidor en 8080
	http.ListenAndServe(":8080", nil)
}
func saludar(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hola mundo")
}

//primer handler(funcion que se encargara de responder aquellas peticiones que nos hacen)
//server mux  .. encargado de enrutar las peticiones
