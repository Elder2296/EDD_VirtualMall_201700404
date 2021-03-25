package Saludar

import (
	"fmt"
)

type nodoarbol struct {
	producto Producto
	altura   int
	izq, der *nodoarbol
}

func newNodo(producto Producto) *nodoarbol {
	return &nodoarbol{producto, 0, nil, nil}
}

type AVL struct {
	raiz *nodoarbol
}

func (avl *AVL) Init() {
	avl.raiz = nil
}

func NewAVL() *AVL {
	return &AVL{nil}
}

func max(val1 int, val2 int) int {
	if val1 > val2 {
		return val1
	}
	return val2
}
func altura(temp *nodoarbol) int {
	if temp != nil {
		return temp.altura
	}
	return -1
}

func rotacionIzquierda(temp **nodoarbol) {
	aux := (*temp).izq
	(*temp).izq = aux.der
	aux.der = *temp
	(*temp).altura = max(altura((*temp).der), altura((*temp).izq)) + 1
	aux.altura = max(altura((*temp).izq), (*temp).altura) + 1
	*temp = aux
}
func rotacionDerecha(temp **nodoarbol) {
	aux := (*temp).der
	(*temp).der = aux.izq
	aux.izq = *temp
	(*temp).altura = max(altura((*temp).der), altura((*temp).izq)) + 1
	aux.altura = max(altura((*temp).der), (*temp).altura) + 1
	*temp = aux
}
func rotacionDobleIzquierda(temp **nodoarbol) {
	rotacionDerecha(&(*temp).izq)
	rotacionIzquierda(temp)
}
func rotacionDobleDerecha(temp **nodoarbol) {
	rotacionIzquierda(&(*temp).der)
	rotacionDerecha(temp)
}
func (avl *AVL) Insertar(producto Producto) {
	insert(producto, &avl.raiz)
}

func insert(producto Producto, root **nodoarbol) {
	if (*root) == nil {
		*root = newNodo(producto)
		//fmt.Println("Agrego algo")
		return
	}
	if producto.Codigo < (*root).producto.Codigo {
		insert(producto, &(*root).izq)
		if (altura((*root).izq) - altura((*root).der)) == -2 {
			if producto.Codigo < (*root).izq.producto.Codigo {
				rotacionIzquierda(root)
			} else {
				rotacionDobleIzquierda(root)
			}
		}
	} else if producto.Codigo > (*root).producto.Codigo {
		insert(producto, &(*root).der)
		if (altura((*root).der) - altura((*root).izq)) == 2 {
			if producto.Codigo > (*root).der.producto.Codigo {
				rotacionDerecha(root)
			} else {
				rotacionDobleDerecha(root)
			}
		}
	} else {
		search(producto, *root).producto.Cantidad += producto.Cantidad

	}

	(*root).altura = max(altura((*root).izq), altura((*root).der)) + 1
}

func search(producto Producto, root *nodoarbol) *nodoarbol {
	aux := root

	for aux != nil {
		if (*aux).producto.Codigo == producto.Codigo {
			return aux
		} else if producto.Codigo < aux.producto.Codigo {
			aux = aux.izq
		} else {
			aux = aux.der

		}

	}
	return aux

}

func (avl *AVL) Print() {
	if avl.raiz == nil {
		fmt.Println("La raiz ha sido inicializada")

	}
	inOrden(avl.raiz)
}

func inOrden(temp *nodoarbol) {
	if temp != nil {
		inOrden(temp.izq)
		fmt.Println("Codigo: ", temp.producto.Codigo, " Nombre: ", temp.producto.Nombre, " stock: ", temp.producto.Cantidad)
		inOrden(temp.der)
	}
}
