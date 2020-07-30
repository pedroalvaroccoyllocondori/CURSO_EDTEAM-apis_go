package almacenamiento

import (
	"fmt"

	"github.com/pedroalvaroccoyllocondori/apis_go/modelo"
)

type Memoria struct {
	actualID int
	Personas map[int]modelo.Persona
}

// funcion constructora
func NuevaMemoria() Memoria {
	personas := make(map[int]modelo.Persona)
	//los mapas necesitan ser inicializados
	return Memoria{
		actualID: 0,
		Personas: personas,
	}

}

//funcion crear
func (referencia *Memoria) Crear(persona *modelo.Persona) error {
	if persona == nil {
		return modelo.ErrorPersonaNoPuedeSerNula
	}
	referencia.actualID++
	referencia.Personas[referencia.actualID] = *persona
	return nil
}

// funcion actualizar

func (referencia *Memoria) Actualizar(ID int, persona *modelo.Persona) error {
	if persona == nil {
		return modelo.ErrorPersonaNoPuedeSerNula
	}
	if _, ok := referencia.Personas[ID]; !ok {
		return fmt.Errorf(" ID: %d: %w", ID, modelo.ErrorIDPersonaNoExiste)

	}

	referencia.Personas[ID] = *persona

	return nil
}

//funcion borrar

func (referencia *Memoria) Quitar(ID int) error {

	if _, ok := referencia.Personas[ID]; !ok {
		return fmt.Errorf("ID: %d: %w", ID, modelo.ErrorIDPersonaNoExiste)
	}
	delete(referencia.Personas, ID)
	return nil
}
func (referencia *Memoria) ObtenerID(ID int) (modelo.Persona, error) {
	persona, ok := referencia.Personas[ID]
	if !ok {
		return persona, fmt.Errorf("ID: %d :%w", ID, modelo.ErrorIDPersonaNoExiste)
	}
	return persona, nil
}

//funcion obtener todas las personas

func (referencia *Memoria) ObternerTodos() (modelo.Personas, error) {
	var resultado modelo.Personas
	for _, valor := range referencia.Personas {
		resultado = append(resultado, valor)
	}
	return resultado, nil
}
