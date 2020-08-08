package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New() // creacion de  instancia de echo
	e.GET("/", saludar)
	e.Start(":8080") //
}

func saludar(c echo.Context) error { // registrar una ruta por el metodo get
	return c.String(http.StatusOK, "Hello, World!") // handler clasico
}
