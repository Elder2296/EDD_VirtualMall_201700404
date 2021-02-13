package saludar

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

type casilla struct {
	Indice       string
	Categoria    string
	Calificacion int
	listatiendas ListaD.ListaDoble
}

func WorkData(datos *Sobre) {
	var nfil int
	var ncol int
	//tam := nfil * ncol * 5
	//var casillas [tam]casilla
	nfil = len(datos.Datos)
	ncol = len(datos.Datos[0].Departamentos)
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

				c++
				fmt.Println("No. de vuelta: ", c)
			}

		}
	}
}
