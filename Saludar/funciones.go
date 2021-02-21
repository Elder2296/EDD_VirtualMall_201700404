package Saludar

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

var Cel *[]Casilla
var bdata *Sobre

func Save() *Sobre {
	return bdata

}

func WorkData(datos *Sobre) {
	bdata = datos
	nfil := len(datos.Datos)
	ncol := len(datos.Datos[0].Departamentos)
	tam := nfil * ncol * 5

	Celdas := make([]Casilla, tam)

	for i := 0; i < nfil; i++ {

		for j := 0; j < ncol; j++ {

			for k := 0; k < 5; k++ {
				Celdas[i+nfil*(j+ncol*k)].Indice = datos.Datos[i].Indice
				Celdas[i+nfil*(j+ncol*k)].Categoria = datos.Datos[i].Departamentos[j].Nombre
				Celdas[i+nfil*(j+ncol*k)].Calificacion = k + 1
				lista := NewList()

				for m := 0; m < len(datos.Datos[i].Departamentos[j].Tiendas); m++ {

					if Celdas[i+nfil*(j+ncol*k)].Calificacion == datos.Datos[i].Departamentos[j].Tiendas[m].Calificacion {
						lista.Insert(datos.Datos[i].Departamentos[j].Tiendas[m])
					}

				}

				Celdas[i+nfil*(j+ncol*k)].listatiendas = *lista

			}

		}
	}
	Cel = &Celdas

	fmt.Println("SE AGREGARON LOS DATOS")

}

var Tienda Store

func GetData(peticion Peticion) Store {

	for _, dato := range *Cel {

		if dato.Categoria == peticion.Departamento && dato.Calificacion == peticion.Calificacion {
			if dato.listatiendas.Encontro(peticion.Nombre) == 1 {
				Tienda = dato.listatiendas.GetTienda(peticion.Nombre)
				break
			}
		}

	}

	return Tienda

}

var hay int

func TestIndex(indice int) int {
	hay = 0
	c := 0
	for _, dato := range *Cel {

		c++
		if c == indice {
			if dato.listatiendas.size != 0 {
				hay = 1

			}

		}
	}
	return hay

}
func GetArrayStore(indice int) []Store {
	c := 0
	for _, dato := range *Cel {
		c++
		if c == indice {
			return dato.listatiendas.GetArray()
		}

	}
	return nil

}

var t int

func DeleteStore(peticion Peticion) int {
	t := 0
	Celdas2 := make([]Casilla, len(*Cel))

	c := 0

	for _, dato := range *Cel {

		if dato.Categoria == peticion.Departamento && dato.Calificacion == peticion.Calificacion && dato.listatiendas.size != 0 {

			if dato.listatiendas.EncontrarTienda(peticion.Nombre) == 1 {
				t = dato.listatiendas.DeleteStore(peticion.Nombre)
				fmt.Println("Eliminacion correcta")
				dato.listatiendas.Print()

			}

		}
		Celdas2[c] = dato
		c++

	}
	Cel = &Celdas2

	//Imprimir()
	return t
}

var direc string
var tiendax, linea3 string

func CreateFile() {
	encabezado := "digraph reporte{\n rankdir=LR"
	var arreglo, arreglo2, cabezas, conexion string
	salto := "\n"
	direc := ""

	arreglo += encabezado + salto
	c := 0
	for _, dato := range *Cel {

		linea := "node [shape=circle, style= filled, label= \""
		cali := strconv.Itoa(dato.Calificacion)

		linea += dato.Categoria + " " + dato.Indice + " " + cali + "\"]" + dato.Indice + cali + dato.Categoria + salto

		if (c + 1) == len(*Cel) {
			direc += dato.Indice + cali + dato.Categoria

		} else {
			direc += dato.Indice + cali + dato.Categoria + " -> "

		}

		if dato.listatiendas.size != 0 {

			for i := 0; i < len(dato.listatiendas.GetArray()); i++ {
				linea2 := "node [shape=box, style= filled, label= \""
				linea2 += dato.listatiendas.GetArray()[i].Nombre + "\"]" + dato.listatiendas.GetArray()[i].Nombre + salto
				arreglo2 += linea2

			}

			linea3 = dato.Indice + cali + dato.Categoria + " -> " + dato.listatiendas.cabeza.store.Nombre + salto
			cabezas += linea3
			conexion += dato.listatiendas.conexiones()

		}

		c++
		arreglo += linea

	}
	arreglo += direc + salto
	arreglo += arreglo2
	arreglo += cabezas
	arreglo += conexion
	arreglo += salto + "}"
	er := ioutil.WriteFile("grafico.dot", []byte(arreglo), 0644)

	if er != nil {
		log.Fatal(er)
	} else {

		fmt.Println("archivo creado con exito!!!")

	}

}

/*func Imprimir() {
	for _, dato := range *Cel {

		fmt.Println(" departamento: ", dato.Categoria, " calificacion ", dato.Calificacion)
		fmt.Println("tiendas: ")
		dato.listatiendas.Print()

	}
}*/
