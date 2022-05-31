package main

import (
	"log"
	"github.com/juanrevello/twittor/handlers"
	"github.com/juanrevello/twittor/bd"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("sin conexion a la BD")
		return
	}
	handlers.Manejadores()
}