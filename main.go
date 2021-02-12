package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http" // para utilizar respuestas y peticiones al servidor

	"github.com/gorilla/mux" // para crear un enrutador
)

type store struct {
	Nombre       string `json:Nombre`
	Descripcion  string `json:Descripcion`
	Contacto     string `json:Contacto`
	Calificacion int    `json: Calificacion`
}

type Departament struct {
	Nombre  string  `json: Nombre`
	Tiendas []store `json: Tiendas`
}

type Data struct {
	Indice        string        `json:Indice`
	Departamentos []Departament `json:Departamentos`
}
type Sobre struct {
	Datos []Data `json:Datos`
}

func rutainit(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Raiz de la rutas masacres")

}

func createData(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var data Sobre

	json.Unmarshal(reqBody, &data)
	fmt.Println(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)

}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	// def of routes
	router.HandleFunc("/", rutainit).Methods("GET")

	router.HandleFunc("/cargartienda", createData).Methods("POST")
	router.HandleFunc("/Human", createHuman).Methods("POST")

	//crear un servidor http
	//http.ListenAndServer(":3000", router)

	log.Fatal(http.ListenAndServe(":3000", router))

}
