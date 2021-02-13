package ListaD

import (
	"fmt"

	saludar "github.com/Elder2296/EDD_VirtualMall_201700404/Saludar"
)

type nodo struct {
	next, back *nodo
	store      saludar.Store
}

type ListaDoble struct {
	cabeza *nodo
	size   int
}

func newList() *ListaDoble { //contructor
	return &ListaDoble{cabeza: nil, size: 0}
}

func (this *ListaDoble) Insert(store saludar.Store) {
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
	fmt.Println("Terminos en la lista: ", this.size)
}
