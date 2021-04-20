package Saludar

type Enlace struct {
	Nombre    string `json:Nombre`
	Distancia int    `json:Distancia`
}

type Nodo struct {
	Nombre  string   `json:Nombre`
	Enlaces []Enlace `json:Enlaces`
}
type Grafo struct {
	Nodos                []Nodo `json:Nodos`
	PosicionInicialRobot string `json:PosicionInicialRobot`
	Entrega              string `json:Entrega`
}
