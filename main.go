package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http" // para utilizar respuestas y peticiones al servidor

	"github.com/gorilla/mux" // para crear un enrutador
)

type store struct {
	Name         string `json:Name`
	Description  string `json:Description`
	Contact      string `json.Contact`
	Calificacion int    `json: Calificacion`
}

type allStores []store

var stores = allStores{
	{
		Name:         "juanito",
		Description:  "tienda de juanito",
		Contact:      "54203744",
		Calificacion: 5,
	},
}

func getStores(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(stores)
}

func rutainit(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Raiz de la rutas masacres")

}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	// def of routes
	router.HandleFunc("/", rutainit)
	router.HandleFunc("/stores", getStores).Methods("GET")

	//crear un servidor http
	//http.ListenAndServer(":3000", router)

	log.Fatal(http.ListenAndServe(":3000", router))

}
