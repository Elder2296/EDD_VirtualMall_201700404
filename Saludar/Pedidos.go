package Saludar

type Item struct {
	Id       string `json:Id`
	Codigo   string `json:Codigo`
	Nombre   string `json:Nombre`
	Cantidad string `json:Cantidad`
	Precio   string `json:Precio`
}
type Pedido struct {
	Total string `json:Total`
	Datos []Item `json:Datos`
}
