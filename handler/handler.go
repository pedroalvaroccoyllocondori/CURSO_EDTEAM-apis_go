package handler

import (
	"github.com/pedroalvaroccoyllocondori/apis_go/modelo"
)

type Almacenamiento interface {
	Crear(persona *modelo.Persona) error
	Actualizar(ID int, persona *modelo.Persona) error
	Quitar(ID int) error
	ObtenerID(ID int) (modelo.Persona, error)
	ObternerTodos() (modelo.Personas, error)
}
