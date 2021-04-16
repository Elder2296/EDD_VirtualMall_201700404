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
	"github.com/rs/cors"
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
func cargarInventario(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	var data saludar.Stock
	json.Unmarshal(reqBody, &data)
	//fmt.Println(data)
	saludar.IngresarInventario(&data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)

}
func getTiendas(w http.ResponseWriter, r *http.Request) {

	tiendas := saludar.GetT_tiendas()
	//fmt.Println(tiendas)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(tiendas)

}
func verProductos(w http.ResponseWriter, r *http.Request) {

	reqBody, _ := ioutil.ReadAll(r.Body)
	//fmt.Println("VAMOS A VER ", reqBody)
	var peti saludar.Peticionproductos

	json.Unmarshal(reqBody, &peti)

	//fmt.Println("Id: ", peti)
	productos := saludar.GetProducts(peti)
	//fmt.Println("Entro a esta peticion")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(productos)

}
func hacerpedido(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	var pedido saludar.Pedido

	json.Unmarshal(reqBody, &pedido)

	fmt.Println("Entro a hacer el pedido")
	fmt.Println(pedido)

	fmt.Println(reqBody)

}
func main() {
	fmt.Println("Entro al sistema")
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", rutainit).Methods("GET")

	router.HandleFunc("/cargartienda", createData).Methods("POST")
	router.HandleFunc("/TiendaEspecifica", getTienda).Methods("POST")
	router.HandleFunc("/id/{numero}", getPosicion).Methods("GET")
	router.HandleFunc("/Eliminar", deleteStore).Methods("DELETE")
	router.HandleFunc("/guardar", saveFile).Methods("GET")
	router.HandleFunc("/getArreglo", getArreglo).Methods("GET")
	router.HandleFunc("/cargarInventario", cargarInventario).Methods("POST")
	router.HandleFunc("/Tiendas", getTiendas).Methods("GET")
	router.HandleFunc("/Verproductos", verProductos).Methods("POST")
	router.HandleFunc("/HacerPedido", hacerpedido).Methods("POST")
	newcors := cors.New(cors.Options{AllowedOrigins: []string{"http://localhost:8001"},
		AllowCredentials: true})

	handler := newcors.Handler(router)
	log.Fatal(http.ListenAndServe(":3000", handler))

	/*mylista := saludar.NewListaSimple()

	producto := saludar.Producto{"s8", 1234, "El smartphone del futuro", 2500.00, 25, "https://i.blogs.es/7a4489/galaxy-s8-4/450_1000.jpg"}
	producto2 := saludar.Producto{"s7", 1237, "El smartphone del futuro", 2500.00, 25, "https://i.blogs.es/7a4489/galaxy-s8-4/450_1000.jpg"}
	producto3 := saludar.Producto{"s6", 1236, "El smartphone del futuro", 2500.00, 25, "https://i.blogs.es/7a4489/galaxy-s8-4/450_1000.jpg"}
	producto4 := saludar.Producto{"s10", 1227, "El smartphone del futuro", 2500.00, 25, "https://i.blogs.es/7a4489/galaxy-s8-4/450_1000.jpg"}
	producto5 := saludar.Producto{"s10", 1227, "El smartphone del futuro", 2500.00, 25, "https://i.blogs.es/7a4489/galaxy-s8-4/450_1000.jpg"}

	mylista.Insertar(producto)
	mylista.Insertar(producto2)
	mylista.Insertar(producto3)
	mylista.Insertar(producto4)
	mylista.Insertar(producto5)
	mylista.Imprimir()
	*/

	/*var arbol saludar.AVL

	miavl := saludar.NewAVL()
	arbol = *miavl
	producto := saludar.Producto{"s8", 1234, "El smartphone del futuro", 2500.00, 25, "https://i.blogs.es/7a4489/galaxy-s8-4/450_1000.jpg"}
	producto2 := saludar.Producto{"s7", 1237, "El smartphone del futuro", 2500.00, 25, "https://i.blogs.es/7a4489/galaxy-s8-4/450_1000.jpg"}
	producto3 := saludar.Producto{"s6", 1236, "El smartphone del futuro", 2500.00, 25, "https://i.blogs.es/7a4489/galaxy-s8-4/450_1000.jpg"}
	producto4 := saludar.Producto{"s10", 1227, "El smartphone del futuro", 2500.00, 25, "https://i.blogs.es/7a4489/galaxy-s8-4/450_1000.jpg"}
	producto5 := saludar.Producto{"s10", 1227, "El smartphone del futuro", 2500.00, 25, "https://i.blogs.es/7a4489/galaxy-s8-4/450_1000.jpg"}
	miavl.Insertar(producto)
	miavl.Insertar(producto2)
	miavl.Insertar(producto3)
	miavl.Insertar(producto4)
	miavl.Insertar(producto5)
	miavl.Print()*/

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
