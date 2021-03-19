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
func (this *ListaDoble) Encontro(name string) int { // solo dice si hay una coincidencia

	aux := this.cabeza
	Encontro := 0
	for aux != nil {

		if aux.store.Nombre == name {
			Encontro = 1
			break
		} else {
			aux = aux.next
		}
	}
	return Encontro
}

func (this *ListaDoble) GetTienda(name string) Store {
	aux := this.cabeza
	for aux != nil {
		if aux.store.Nombre == name {
			break
		} else {
			aux = aux.next
		}
	}
	return aux.store
}

var r int

func (this *ListaDoble) EncontrarTienda(nombre string) int {
	r = 0
	aux := this.cabeza

	if aux != nil {
		for aux != nil {
			if aux.store.Nombre == nombre {
				r = 1
				break
			} else {
				aux = aux.next
			}
		}

	}

	return r
}
func (this *ListaDoble) GetArray() []Store {

	array := make([]Store, this.size)
	pos := 0
	aux := this.cabeza
	for aux != nil {
		array[pos] = aux.store
		aux = aux.next
		pos++

	}
	return array

}

var cumplio int

func (this *ListaDoble) DeleteStore(nombre string) int {
	aux := this.cabeza
	cumplio = 0

	if this.size == 1 {
		this.cabeza = nil
		this.size--
		cumplio = 1
	} else {

		if this.cabeza.store.Nombre == nombre {
			this.cabeza = aux.next
			this.cabeza.back = nil
			cumplio = 1
			this.size--

		} else {
			for aux != nil {

				if aux.store.Nombre == nombre {
					temp := aux.back
					temp.next = aux.next
					cumplio = 1
					this.size--
					break

				}
				aux = aux.next
			}

		}

	}

	return cumplio
}
func (this *ListaDoble) conexiones() string {
	aux := this.cabeza
	salida := ""
	for aux != nil {
		if aux.next != nil {
			salida += aux.store.identi + " -> " + aux.next.store.identi + "\n"
			salida += aux.next.store.identi + " -> " + aux.store.identi + "\n"

		}
		aux = aux.next

	}
	return salida

}
