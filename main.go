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
	// grupo de rutas
	//e.POST("/personas/crear",crear)
	//e.GET("/personas/consultar",consultar)
	//e.PUT("/personas/actualizar",actualizar)
	//e.DELETE("/personas/borrar",borrar)

	personas := e.Group("/personas")
	personas.POST("", crear)
	personas.GET("/:id", consultar)
	personas.PUT("/:id", actualizar)
	personas.DELETE("/:id", borrar)

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

func crear(c echo.Context) error {
	return c.String(http.StatusOK, "creado")
}
func consultar(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "consultado"+id)
}
func actualizar(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "actualizado"+id)
}
func borrar(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "borrar"+id)
}
