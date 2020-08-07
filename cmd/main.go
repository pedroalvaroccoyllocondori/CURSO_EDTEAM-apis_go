package main

import (
	"log"
	"net/http"

	"github.com/pedroalvaroccoyllocondori/apis_go/autorizacion"

	"github.com/pedroalvaroccoyllocondori/apis_go/almacenamiento"
	"github.com/pedroalvaroccoyllocondori/apis_go/handler"
)

func main() {

	err := autorizacion.CargarArchivos("certificates/app.rsa", "certificates/app.rsa.pub")
	if err != nil {
		log.Fatalf("no se pudo cargar los certificados: %v", err)
	}

	almacen := almacenamiento.NuevaMemoria()
	mux := http.NewServeMux()

	handler.RutaPersona(mux, &almacen)
	handler.RutaLogin(mux, &almacen)
	log.Printf("servidor iniciado en el puerto 8080")
	err = http.ListenAndServe(":8080", mux) //funcion que devuelve un error
	if err != nil {
		log.Printf("error en el servidor:%v \n", err)
	}

}
