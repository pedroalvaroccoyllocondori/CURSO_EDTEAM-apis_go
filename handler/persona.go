package handler

import (
	"encoding/json"
	"errors"
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

		respuesta := nuevaRespuesta(Error, "metodo no permitido", nil)
		respuestaJSON(w, http.StatusBadRequest, respuesta)
		return
	}

	data := modelo.Persona{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {

		respuesta := nuevaRespuesta(Mensaje, "la persona no tiene la estructura correcta", nil)
		respuestaJSON(w, http.StatusBadRequest, respuesta)

		return
	}

	err = p.almacenamiento.Crear(&data)
	if err != nil {

		respuesta := nuevaRespuesta(Error, "hubo  un problema al crear la persona", nil)
		respuestaJSON(w, http.StatusBadRequest, respuesta)

		return
	}

	respuesta := nuevaRespuesta(Mensaje, "persona creada correctamente", nil)
	respuestaJSON(w, http.StatusBadRequest, respuesta)

}

func (p *persona) actualizar(w http.ResponseWriter, r *http.Request) {
	//metodo   para validar el codigo de respuesta
	if r.Method != http.MethodPut {

		respuesta := nuevaRespuesta(Error, "Metodo no permitido", nil)
		respuestaJSON(w, http.StatusBadRequest, respuesta)
		return
	}
	// capturar los parametros
	ID, err := strconv.Atoi(r.URL.Query().Get("id")) //convetira aun numero// devuelve el string con su valor
	//metodo para validar el error
	if err != nil {

		respuesta := nuevaRespuesta(Error, "el id debe ser un numero entero positivo", nil)
		respuestaJSON(w, http.StatusBadRequest, respuesta)

		return
	}
	// validar si la persona esta creada
	data := modelo.Persona{}
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {

		respuesta := nuevaRespuesta(Error, "la persona no tiene la estructura correcta", nil)
		respuestaJSON(w, http.StatusBadRequest, respuesta)

		return
	}

	err = p.almacenamiento.Actualizar(ID, &data)
	if err != nil {

		respuesta := nuevaRespuesta(Error, "ok", nil)
		respuestaJSON(w, http.StatusInternalServerError, respuesta)

		return
	}

	respuesta := nuevaRespuesta(Mensaje, "persona actualizada correctamente", nil)
	respuestaJSON(w, http.StatusOK, respuesta)

}

func (p *persona) quitar(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {

		respuesta := nuevaRespuesta(Error, "Metodo no permitido", nil)
		respuestaJSON(w, http.StatusBadRequest, respuesta)

		return
	}
	ID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {

		respuesta := nuevaRespuesta(Error, "el id debe ser un numero entero positivo", nil)
		respuestaJSON(w, http.StatusBadRequest, respuesta)

		return
	}

	err = p.almacenamiento.Quitar(ID)
	if errors.Is(err, modelo.ErrorIDPersonaNoExiste) {
		respuesta := nuevaRespuesta(Error, "el id de la peersona no existe registado el el host", nil)
		respuestaJSON(w, http.StatusBadRequest, respuesta)
		return
	}

	if err != nil {
		respuesta := nuevaRespuesta(Error, "ocurrio un error al eliminar el registro", nil)
		respuestaJSON(w, http.StatusInternalServerError, respuesta)

		return

	}

	respuesta := nuevaRespuesta(Mensaje, "ok", nil)
	respuestaJSON(w, http.StatusOK, respuesta)

}

func (p *persona) obtenerID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respuesta := nuevaRespuesta(Error, "Método no permitido", nil)
		respuestaJSON(w, http.StatusBadRequest, respuesta)
		return
	}

	ID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {

		respuesta := nuevaRespuesta(Error, "El id debe ser un número entero positivo", nil)
		respuestaJSON(w, http.StatusBadRequest, respuesta)
		return
	}

	data, err := p.almacenamiento.ObtenerID(ID)
	if errors.Is(err, modelo.ErrorIDPersonaNoExiste) {

		respuesta := nuevaRespuesta(Error, "El ID de la persona no existe", nil)
		respuestaJSON(w, http.StatusBadRequest, respuesta)

		return
	}
	if err != nil {

		respuesta := nuevaRespuesta(Error, "Ocurrió un error al elminar el registro", nil)
		respuestaJSON(w, http.StatusInternalServerError, respuesta)
		return
	}

	respuesta := nuevaRespuesta(Mensaje, "ok", data)
	respuestaJSON(w, http.StatusOK, respuesta)
}

func (p *persona) obtenerTodos(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {

		respuesta := nuevaRespuesta(Error, "Metodo no permitido", nil)
		respuestaJSON(w, http.StatusBadRequest, respuesta)

		return
	}

	data, err := p.almacenamiento.ObternerTodos()
	if err != nil {

		respuesta := nuevaRespuesta(Error, "hubo un error al obtener todas las personas", nil)
		respuestaJSON(w, http.StatusOK, respuesta)

		return
	}

	respuesta := nuevaRespuesta(Mensaje, "ok", data)
	respuestaJSON(w, http.StatusOK, respuesta)

}
