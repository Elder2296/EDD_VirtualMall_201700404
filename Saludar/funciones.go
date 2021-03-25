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

func IngresarInventario(stoke *Stock) {
	/*ntiendas := len(stoke.Inventarios)
	fmt.Println("Tiendas que seran afectadas: ", ntiendas)*/

	Celdas2 := make([]Casilla, len(*Cel))
	c := 0
	ninven := 0
	/*for _, elemento := range stoke.Inventarios {
		fmt.Println("Tienda: ", elemento.Tienda, " Departamento: ", elemento.Departamento, " Calificacion: ", elemento.Calificacion)

	}*/

	for _, dato := range *Cel {

		for _, elemento := range stoke.Inventarios {
			if dato.Categoria == elemento.Departamento && dato.Calificacion == elemento.Calificacion && dato.listatiendas.Encontro(elemento.Tienda) != 0 {

				//dato.listatiendas.GetTienda(elemento.Tienda).Arbol.Print()
				for _, objeto := range elemento.Productos {
					dato.listatiendas.GetTienda(elemento.Tienda).Arbol.Insertar(objeto)
					ninven++
				}

				/*dato.listatiendas.GetTienda(elemento.Tienda).Arbol.Print()
				fmt.Println("------------------------------")**/
			}

		}
		Celdas2[c] = dato
		c++

	}
	Cel = &Celdas2

	fmt.Println("Se hicieorn  ", ninven, " inserciones en los arboles")

}

func WorkData(datos *Sobre) {
	bdata = datos
	nfil := len(datos.Datos)
	ncol := len(datos.Datos[0].Departamentos)
	tam := nfil * ncol * 5

	cont := 1

	Celdas := make([]Casilla, tam)

	cont2 := 0

	for i := 0; i < nfil; i++ {

		for j := 0; j < ncol; j++ {

			for k := 0; k < 5; k++ {

				Celdas[i+nfil*(j+ncol*k)].Indice = datos.Datos[i].Indice
				Celdas[i+nfil*(j+ncol*k)].Categoria = datos.Datos[i].Departamentos[j].Nombre
				Celdas[i+nfil*(j+ncol*k)].Calificacion = k + 1
				Celdas[i+nfil*(j+ncol*k)].identi = "cas" + strconv.Itoa(cont)

				lista := NewList()

				for m := 0; m < len(datos.Datos[i].Departamentos[j].Tiendas); m++ {

					if Celdas[i+nfil*(j+ncol*k)].Calificacion == datos.Datos[i].Departamentos[j].Tiendas[m].Calificacion {
						datos.Datos[i].Departamentos[j].Tiendas[m].identi = "tie" + strconv.Itoa(cont2)
						datos.Datos[i].Departamentos[j].Tiendas[m].Arbol = NewAVL()
						//fmt.Println("----------------------------------------------")
						//producto := Producto{"s8", 1234, "El smartphone del futuro", 2500.00, 25, "https://i.blogs.es/7a4489/galaxy-s8-4/450_1000.jpg"}
						//datos.Datos[i].Departamentos[j].Tiendas[m].Arbol.Insertar(producto)
						//datos.Datos[i].Departamentos[j].Tiendas[m].Arbol.Print()
						lista.Insert(datos.Datos[i].Departamentos[j].Tiendas[m])
						cont2++
					}

				}
				cont++

				Celdas[i+nfil*(j+ncol*k)].listatiendas = *lista

				cont++

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
				//dato.listatiendas.GetTienda(peticion.Nombre).Arbol.Print()
				Tienda = dato.listatiendas.GetTienda(peticion.Nombre)
				Tienda.Arbol = dato.listatiendas.GetTienda(peticion.Nombre).Arbol

				break
			}
		}

	}
	Tienda.Arbol.Print()
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

		linea += dato.Categoria + " " + dato.Indice + " " + cali + "\"]" + dato.identi + salto

		if (c + 1) == len(*Cel) {
			direc += dato.identi

		} else {
			direc += dato.identi + " -> "

		}

		if dato.listatiendas.size != 0 {

			for i := 0; i < len(dato.listatiendas.GetArray()); i++ {
				linea2 := "node [shape=box, style= filled, label= \""
				linea2 += dato.listatiendas.GetArray()[i].Nombre + "\"]" + dato.listatiendas.GetArray()[i].identi + salto
				arreglo2 += linea2

			}

			linea3 = dato.identi + " -> " + dato.listatiendas.cabeza.store.identi + salto
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
