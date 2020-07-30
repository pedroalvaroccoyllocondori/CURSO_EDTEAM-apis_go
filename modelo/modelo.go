package modelo

import "errors"

var (
	//la persona no puede ser nula
	ErrorPersonaNoPuedeSerNula = errors.New("la persona no puede ser nula")
	//persona no existe
	ErrorIDPersonaNoExiste = errors.New("la persona no existe")
)
