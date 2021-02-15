package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http" // para utilizar respuestas y peticiones al servidor

	saludar "github.com/Elder2296/EDD_VirtualMall_201700404/Saludar"
	"github.com/gorilla/mux"
	// para crear un enrutador
)

func rutainit(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Raiz de la rutas masacres")

}

func createData(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var data saludar.Sobre

	json.Unmarshal(reqBody, &data)

	//fmt.Println(data)
	saludar.WorkData(&data)

	/*for _, dato := range data.Datos {
		fmt.Println(dato.Indice)

	}*/

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)

}

var Tienda saludar.Store

func getTienda(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var peti saludar.Peticion
	json.Unmarshal(reqBody, &peti)

	Tienda = saludar.GetData(peti)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(Tienda)

}

func main() {

	router := mux.NewRouter().StrictSlash(true)
	// def of routes
	router.HandleFunc("/", rutainit).Methods("GET")

	router.HandleFunc("/cargartienda", createData).Methods("POST")
	router.HandleFunc("/TiendaEspecifica", getTienda).Methods("GET")
	//router.HandleFunc("/getArreglo", getArray).Methods("GET")//REPORTE DE GRAPHVIZ

	log.Fatal(http.ListenAndServe(":3000", router))
	/*list := new(ListaD.ListaDoble)

	store1 := new(saludar.Store)
	store1.Nombre = "tienda1"
	store2 := new(saludar.Store)
	store2.Nombre = "tienda2"
	store3 := new(saludar.Store)
	store3.Nombre = "tienda3"

	list.Insert(*store1)
	list.Insert(*store2)
	list.Insert(*store3)
	list.Print()*/
}
