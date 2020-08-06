package handler

import (
	"net/http"

	"github.com/pedroalvaroccoyllocondori/apis_go/middleware"
)

func RutaPersona(mux *http.ServeMux, almacenamiento Almacenamiento) {
	h := nuevaPersona(almacenamiento)

	mux.HandleFunc("/v1/personas/crear", middleware.Log(middleware.Autentificacion(h.crear)))
	mux.HandleFunc("/v1/personas/actualizar", h.actualizar)
	mux.HandleFunc("/v1/personas/quitar", middleware.Log(h.quitar))
	mux.HandleFunc("/v1/personas/obtenerId", h.obtenerID)
	mux.HandleFunc("/v1/personas/obtener-todos", middleware.Log(h.obtenerTodos))

}
