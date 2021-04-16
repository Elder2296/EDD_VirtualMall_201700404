package Saludar

import "fmt"

type NodoSimple struct {
	producto  Producto
	siguiente *NodoSimple
}

type ListaSimple struct {
	raiz *NodoSimple
	size int
}

func NewListaSimple() *ListaSimple {
	return &ListaSimple{raiz: nil, size: 0}
}
func (this *ListaSimple) Insertar(producto Producto) {
	nuevo := &NodoSimple{producto: producto, siguiente: nil}

	if this.raiz == nil {
		this.raiz = nuevo
	} else {
		aux := this.raiz
		this.raiz = nuevo
		this.raiz.siguiente = aux
	}

	this.size++

}
func (this *ListaSimple) Imprimir() {
	aux := this.raiz
	fmt.Println("-------------- IMPRESION DESDE LA LISTA--------------------")
	for aux != nil {
		fmt.Println("Nombre de Producto: ", aux.producto.Nombre)
		aux = aux.siguiente

	}
}

func (this *ListaSimple) getArreglo() []Producto {
	productos := make([]Producto, this.size)

	aux := this.raiz
	pos := 0
	for aux != nil {
		productos[pos] = aux.producto
		aux = aux.siguiente
		pos++
	}

	return productos

}
