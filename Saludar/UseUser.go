package Saludar

import (
	"fmt"
	"strconv"
)

var listaUsuarios *ListaUsuarios

func InitUsers(usuarios *Usuarios) {
	fmt.Println(usuarios)
	listaUsuarios = NewListaUsuario()
	for _, usuario := range usuarios.Usuarios {
		listaUsuarios.Insert(usuario)
	}
	fmt.Println("Agregado con exito")
}
func InsertnewUser(user Usuario) {
	listaUsuarios.Insert(user)
	fmt.Println("Registro correcto")
}

func Login(datos *Init) *Usuario {

	dpi, _ := strconv.Atoi(datos.Dpi)
	pass := datos.Password

	if listaUsuarios.UserExist(dpi) {
		if listaUsuarios.getUser(dpi).Password == pass {
			return listaUsuarios.getUser(dpi)
		}
	}
	user := new(Usuario)
	user.Cuenta = ""

	return user

}
