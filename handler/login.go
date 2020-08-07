package handler

import (
	"encoding/json"
	"net/http"

	"github.com/pedroalvaroccoyllocondori/apis_go/autorizacion"

	"github.com/pedroalvaroccoyllocondori/apis_go/modelo"
)

type login struct {
	almacenamiento Almacenamiento
}

func NuevoLogin(a Almacenamiento) login {
	return login{a}
}

func (l *login) login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {

		respuesta := nuevaRespuesta(Error, "metodo no permitido", nil)
		respuestaJSON(w, http.StatusBadRequest, respuesta)
		return
	}

	data := modelo.Login{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		respuesta := nuevaRespuesta(Error, "estructura no valida", nil)
		respuestaJSON(w, http.StatusBadRequest, respuesta)
		return
	}
	if !esValidoLogin(&data) {
		respuesta := nuevaRespuesta(Error, "usuario y comntraseña no valida", nil)
		respuestaJSON(w, http.StatusBadRequest, respuesta)
		return
	}

	token, err := autorizacion.GenerarToken(&data)
	if err != nil {
		respuesta := nuevaRespuesta(Error, "no se pudo generar el token", nil)
		respuestaJSON(w, http.StatusInternalServerError, respuesta)
		return
	}

	datatoken := map[string]string{"token": token}
	respuesta := nuevaRespuesta(Mensaje, "ok", datatoken)
	respuestaJSON(w, http.StatusOK, respuesta)

}
func esValidoLogin(data *modelo.Login) bool {
	return data.Email == "contacto@icasoft.team" && data.Contraseña == "12345"
}
