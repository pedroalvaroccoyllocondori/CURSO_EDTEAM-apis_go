package handler

import (
	"encoding/json"
	"net/http"

	"github.com/pedroalvaroccoyllocondori/apis_go/modelo"
)

type persona struct {
	almacenamiento Almacenamiento
}

func nuevaPersona(almacenamiento Almacenamiento) persona {
	return persona{almacenamiento}
}

func (p *persona) crear(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message_type":"error","message":"Metodo no permitido"}`))
		return
	}
	data := modelo.Persona{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message_type":"error","message":"la persona no tiene la estructura correcta"}`))
		return
	}

	err = p.almacenamiento.Crear(&data)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message_type":"error","message":"hubo  un problema al crear la persona"}`))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message_type":"message","message":"persona creada correctamente"}`))
	return

}
