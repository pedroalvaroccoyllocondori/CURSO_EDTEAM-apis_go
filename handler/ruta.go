package handler

import (
	"net/http"
)

func RutaPersona(mux *http.ServeMux, almacenamiento Almacenamiento) {
	h := nuevaPersona(almacenamiento)

	mux.HandleFunc("/v1/personas/crear", h.crear)

}
