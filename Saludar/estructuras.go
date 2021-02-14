package Saludar

import (
	"fmt"

	"github.com/Elder2296/EDD_VirtualMall_201700404/ListaD"
)

type Store struct {
	Nombre       string `json:Nombre`
	Descripcion  string `json:Descripcion`
	Contacto     string `json.Contacto`
	Calificacion int    `json.Calificacion`
}

type Departament struct {
	Nombre  string  `json:Nombre`
	Tiendas []Store `json:Tiendas`
}

type Data struct {
	Indice        string        `json:Indice`
	Departamentos []Departament `json:Departamentos`
}
type Sobre struct {
	Datos []Data `json: Datos`
}

type Casilla struct {
	Indice       string
	Categoria    string
	Calificacion int
	listatiendas ListaD.ListaDoble
}

func WorkData(datos *Sobre) {

	nfil := len(datos.Datos)
	ncol := len(datos.Datos[0].Departamentos)
	tam := nfil * ncol * 5
	celdas := make([]Casilla, tam)
	fmt.Println("tamanio de celdas ", len(celdas))

	c := 0
	//fmt.Println("cuantas indices viene: ", len(datos.Datos))
	for i := 0; i < nfil; i++ {
		//fmt.Println(dato.Indice)
		//fmt.Println("el indice ", dato.Indice, " trae  ", len(dato.Departamentos), " departamentos ")
		for j := 0; j < ncol; j++ {
			/*fmt.Println(dep.Nombre)
			fmt.Println("el departamento ", dep.Nombre, " trae ", len(dep.Tiendas), " tiendas ")
			for _, store := range dep.Tiendas {
				fmt.Println(store.Nombre)
			}*/

			for k := 0; k < 5; k++ {
				/*celdas[i+nfil*(j+ncol*k)].Indice = datos.Datos[i].Indice
				fmt.Println(celdas[i+nfil*(j+ncol*k)].Indice)*/
				c++
				fmt.Println("No. de vuelta: ", c)
			}

		}
	}

	//fmt.Println("tamaño de linealizacion: ", len(celdas))
}
