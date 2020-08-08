package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New() // creacion de  instancia de echo
	e.Use(middleware.Recover())
	e.GET("/", saludar)
	e.GET("/dividir", dividir)
	e.Start(":8080") //
}

func saludar(c echo.Context) error { // registrar una ruta por el metodo get
	return c.String(http.StatusOK, "Hello, World!") // handler clasico
}
func dividir(c echo.Context) error {
	d := c.QueryParam("d")
	f, _ := strconv.Atoi(d)
	if f == 0 {
		return c.String(http.StatusBadRequest, "el valor no puede ser cero")

	}
	r := 3000 / f
	return c.String(http.StatusOK, strconv.Itoa(r))
}
