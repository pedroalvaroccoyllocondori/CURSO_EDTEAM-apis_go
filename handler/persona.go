package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

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

}

func (p *persona) actualizar(w http.ResponseWriter, r *http.Request) {
	//metodo   para validar el codigo de respuesta
	if r.Method != http.MethodPut {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message_type":"error","message":"Metodo no permitido"}`))
		return
	}
	// capturar los parametros
	ID, err := strconv.Atoi(r.URL.Query().Get("id")) //convetira aun numero// devuelve el string con su valor
	//metodo para validar el error
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message_type":error,"message":"el id debe ser un numero entero positivo"}`))
		return
	}
	// validar si la persona esta creada
	data := modelo.Persona{}
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message_type":"error","message":"la persona no tiene la estructura correcta"}`))
		return
	}

	err = p.almacenamiento.Actualizar(ID, &data)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message_type":"message","message":"ok"}`))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message_type":"message","message":"persona actualizada correctamente"}`))

}

func (p *persona) obtenerTodos(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message_type":"error","message":"Metodo no permitido"}`))
		return
	}

	respuesta, err := p.almacenamiento.ObternerTodos()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message_type":error,"message":"hubo un error al obtener todas las personas"}`))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&respuesta)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message_type":"message","message":"ok"}`))
		return
	}

}
