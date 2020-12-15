package main

import (
	"log"

	"github.com/juang77/Gabotor/bd"
	"github.com/juang77/Gabotor/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Son conexion a la BD")
	}
	handlers.Manejadores()
}
