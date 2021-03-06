package handlers

import (
	"github.com/gorilla/mux"
	"github.com/juanrevello/twittor/middlew"
	"github.com/juanrevello/twittor/routers"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

/*Manejadores seteo mi puerto, el handler, y pongo a escuchar al servidor */
func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middlew.ChequeoBD(routers.Login)).Methods("POST")
	router.HandleFunc("/verperfil", middlew.ChequeoBD(middlew.ValidoJWT(routers.VerPerfil))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
