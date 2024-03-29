package Saludar

type Peticion struct {
	Departamento string `json:Departamento`
	Nombre       string `json:Nombre`
	Calificacion int    `json:Calificacion`
}
type Peticionproductos struct {
	Id string `json:Id`
}
type Store struct {
	Nombre       string `json:Nombre`
	Descripcion  string `json:Descripcion`
	Contacto     string `json:Contacto`
	Calificacion int    `json:Calificacion`
	Logo         string `json:Logo`
	Id           string `json:Id`
	Arbol        *AVL
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

type Inventario struct {
	Tienda       string     `json:Tienda`
	Departamento string     `json:Departamento`
	Calificacion int        `json:Calificacion`
	Productos    []Producto `json:Productos`
}
type Stock struct {
	Inventarios []Inventario `json:Inventarios`
}

type Casilla struct {
	Indice       string `json: Indice`
	Departamento string `json:Departamento`
	Calificacion int    `json:Calificacion`
	listatiendas ListaDoble
	identi       string
}
type Producto struct {
	Nombre      string  `json:Nombre`
	Codigo      int     `json:Codigo`
	Descripcion string  `json:Descripcion`
	Precio      float64 `json:Precio`
	Cantidad    int     `json:Cantidad`
	Imagen      string  `json:Imagen`
}

/*
var Cel *[]Casilla

func WorkData(datos *Sobre) {

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

	Imprimir()
	return t
}

func Imprimir() {
	for _, dato := range *Cel {

		fmt.Println(" departamento: ", dato.Categoria, " calificacion ", dato.Calificacion)
		fmt.Println("tiendas: ")
		dato.listatiendas.Print()

	}
}
*/
