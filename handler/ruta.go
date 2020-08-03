package handler

import (
	"net/http"
)

func RutaPersona(mux *http.ServeMux, almacenamiento Almacenamiento) {
	h := nuevaPersona(almacenamiento)

	mux.HandleFunc("/v1/personas/crear", h.crear)
	mux.HandleFunc("/v1/personas/actualizar", h.actualizar)
	mux.HandleFunc("/v1/personas/quitar", h.quitar)
	mux.HandleFunc("/v1/personas/Obtener-todos", h.obtenerTodos)

}
