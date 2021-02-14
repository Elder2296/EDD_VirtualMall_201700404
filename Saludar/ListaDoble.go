package Saludar

import (
	"fmt"
)

type nodo struct {
	next, back *nodo
	store      Store
}

type ListaDoble struct {
	cabeza *nodo
	size   int
}

func NewList() *ListaDoble { //contructor
	return &ListaDoble{cabeza: nil, size: 0}
}

func (this *ListaDoble) Insert(store Store) {
	nuevo := &nodo{nil, nil, store}
	if this.cabeza == nil {
		this.cabeza = nuevo

	} else {
		aux := this.cabeza

		this.cabeza = nuevo
		this.cabeza.next = aux
		aux.back = this.cabeza

	}
	this.size++

}

func (this *ListaDoble) Print() {
	aux := this.cabeza

	for aux != nil {

		fmt.Println(aux.store.Nombre)

		aux = aux.next
	}
	//fmt.Println("Terminos en la lista: ", this.size)
}
