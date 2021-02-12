package saludar

type store struct {
	Nombre       string `json:Nombre`
	Descripcion  string `json:Descripcion`
	Contacto     string `json.Contacto`
	Calificacion int    `json.Calificacion`
}

type Departament struct {
	Nombre  string  `json:Nombre`
	Tiendas []store `json:Tiendas`
}

type Data struct {
	Indice        string        `json:Indice`
	Departamentos []Departament `json:Departamentos`
}
type Sobre struct {
	Datos []Data `json: Datos`
}
