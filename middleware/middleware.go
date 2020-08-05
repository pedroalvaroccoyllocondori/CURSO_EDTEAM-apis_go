package middleware

import (
	"log"
	"net/http"
)

func Log(funcion func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("peticion %q ,metodo %q", r.URL.Path, r.Method)
		funcion(w, r)
	}
}
