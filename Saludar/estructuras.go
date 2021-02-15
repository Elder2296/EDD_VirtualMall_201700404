package Saludar

import "fmt"

type Peticion struct {
	Departamento string `json:Departamento`
	Nombre       string `json:Nombre`
	Calificacion int    `json:Calificacion`
}
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
	Indice       string `json: Indice`
	Categoria    string `json:Categoria`
	Calificacion int    `json:Calificacion`
	listatiendas ListaDoble
}

var Cel *[]Casilla

func WorkData(datos *Sobre) {

	nfil := len(datos.Datos)
	ncol := len(datos.Datos[0].Departamentos)
	tam := nfil * ncol * 5

	Celdas := make([]Casilla, tam)
	//fmt.Println("tamanio de celdas ", len(celdas))

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

	/*for _, dato := range Celdas {

		fmt.Println("Indice: ", dato.Indice, " departamento: ", dato.Categoria, " calificacion ", dato.Calificacion)
		fmt.Println("tiendas: ")
		dato.listatiendas.Print()

	}*/

	fmt.Println("se agregaron los datos")

}

var Tienda Store

func GetData(peticion Peticion) Store {

	/*fmt.Println("Departamento: ", peticion.Departamento)
	fmt.Println("Nombre: ", peticion.Nombre)
	fmt.Println("Calificacion: ", peticion.Calificacion)*/

	for _, dato := range *Cel {

		if dato.Categoria == peticion.Departamento && dato.Calificacion == peticion.Calificacion {
			if dato.listatiendas.Encontro(peticion.Nombre) == 1 {
				Tienda = dato.listatiendas.GetTienda(peticion.Nombre)
				break
			}
		}

		/*fmt.Println("Indice: ", dato.Indice, " departamento: ", dato.Categoria, " calificacion ", dato.Calificacion)
		fmt.Println("tiendas: ")
		dato.listatiendas.Print()*/

	}

	return Tienda

}
