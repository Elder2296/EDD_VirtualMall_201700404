package Saludar

func GetProducts(request Peticionproductos) []Producto {
	//fmt.Println("ESTO SE MUESTRA ", request.Id)
	for _, dato := range *Cel {
		if dato.listatiendas.size != 0 {

			if dato.listatiendas.GetStore_by_Id(request.Id) != 0 {

				return dato.listatiendas.getStore_by_Id(request.Id).Arbol.lista.getArreglo()

			}

		}
	}

	return nil

}
