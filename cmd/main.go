package main

import (
	"log"
	"net/http"

	"github.com/pedroalvaroccoyllocondori/apis_go/almacenamiento"
	"github.com/pedroalvaroccoyllocondori/apis_go/handler"
)

func main() {
	almacen := almacenamiento.NuevaMemoria()
	mux := http.NewServeMux()

	handler.RutaPersona(mux, &almacen)
	log.Printf("servidor iniciado en el puerto 8080")
	err := http.ListenAndServe(":8080", mux) //funcion que devuelve un error
	if err != nil {
		log.Printf("error en el servidor:%v \n", err)
	}

}
