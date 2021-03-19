package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http" // para utilizar respuestas y peticiones al servidor
	"strconv"

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

func getPosicion(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	index, _ := strconv.Atoi(vars["numero"])

	if saludar.TestIndex(index) == 0 {
		mensaje := "no hay registro de tiendas en dicho indice"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(mensaje)

	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(saludar.GetArrayStore(index))

	}

}
func getTienda(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var peti saludar.Peticion
	json.Unmarshal(reqBody, &peti)

	Tienda = saludar.GetData(peti)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(Tienda)

}
func deleteStore(w http.ResponseWriter, r *http.Request) {

	reqBody, _ := ioutil.ReadAll(r.Body)
	var peti saludar.Peticion

	json.Unmarshal(reqBody, &peti)

	fmt.Println("nombre: ", peti.Nombre, " departamento ", peti.Departamento, " calificacion: ", peti.Calificacion)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	//fmt.Println("valor que retorna: ", saludar.DeleteStore(peti))

	if saludar.DeleteStore(peti) == 1 {
		mensaje := "Eliminacion exitosa"

		json.NewEncoder(w).Encode(mensaje)

	} else {
		mensaje := "Delete falled"

		json.NewEncoder(w).Encode(mensaje)

	}

}
func saveFile(w http.ResponseWriter, r *http.Request) {

	arreglo, _ := json.Marshal(saludar.Save())

	err := ioutil.WriteFile("salida.json", arreglo, 0644)

	if err != nil {
		log.Fatal(err)

	} else {
		json.NewEncoder(w).Encode("Se ha guardado exitosamente")
	}

}
func getArreglo(w http.ResponseWriter, r *http.Request) {
	saludar.CreateFile()

}

func main() {

	router := mux.NewRouter().StrictSlash(true)
	// def of route
	router.HandleFunc("/", rutainit).Methods("GET")

	router.HandleFunc("/cargartienda", createData).Methods("POST")
	router.HandleFunc("/TiendaEspecifica", getTienda).Methods("POST")
	router.HandleFunc("/id/{numero}", getPosicion).Methods("GET")
	router.HandleFunc("/Eliminar", deleteStore).Methods("DELETE")
	router.HandleFunc("/guardar", saveFile).Methods("GET")
	router.HandleFunc("/getArreglo", getArreglo).Methods("GET")

	//router.HandleFunc("/getArreglo", getArray).Methods("GET")//REPORTE DE GRAPHVIZ

	log.Fatal(http.ListenAndServe(":3000", router))

	/*list := new(saludar.ListaDoble)

	store1 := new(saludar.Store)
	store1.Nombre = "tienda1"
	store2 := new(saludar.Store)
	store2.Nombre = "tienda2"
	//store3 := new(saludar.Store)
	//store3.Nombre = "tienda3"

	list.Insert(*store1)
	list.Insert(*store2)
	//list.Insert(*store3)
	list.Print()
	list.DeleteStore("tienda2")
	fmt.Println()
	list.Print()

	fmt.Println("")
	list.DeleteStore("tienda1")
	fmt.Println()
	list.Print()*/
	//createfile()

}

func createfile() {
	var cadena string
	linea1 := "este es" + "\n"
	linea2 := "un archivo" + "\n"
	linea3 := "con multi" + "\n"
	linea4 := "lineas" + "\n"
	cadena = linea1 + linea2 + linea3 + linea4

	er := ioutil.WriteFile("prueba.dot", []byte(cadena), 0644)

	if er != nil {
		log.Fatal(er)
	}

}
