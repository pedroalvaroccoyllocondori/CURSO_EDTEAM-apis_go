package handler

import (
	"encoding/json"
	"net/http"
)

const (
	Error   = "error"
	Mensaje = "mensaje"
)

type respuesta struct {
	TipoMensaje string      `json:"message_type"`
	Mensaje     string      `json:"message"`
	Data        interface{} `json:"data"`
}

func nuevaRespuesta(tipoMensaje string, mensaje string, data interface{}) respuesta {
	return respuesta{
		tipoMensaje,
		mensaje,
		data,
	}

}

func respuestaJSON(w http.ResponseWriter, statusCode int, resp respuesta) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(&resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}
}
