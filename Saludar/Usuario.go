package Saludar

type Usuario struct {
	Dpi      int    `json:Dpi`
	Nombre   string `json:Nombre`
	Correo   string `json:Correo`
	Password string `json:Password`
	Cuenta   string `json:Cuenta`
}
type Init struct {
	Dpi      string `json:Dpi`
	Password string `json:Password`
}

type Usuarios struct {
	Usuarios []Usuario `json:Usuarios`
}

type nodoUsuario struct {
	usuario   Usuario
	siguiente *nodoUsuario
}

type ListaUsuarios struct {
	raiz *nodoUsuario
	size int
}

func NewListaUsuario() *ListaUsuarios {
	return &ListaUsuarios{raiz: nil, size: 0}

}
func (this *ListaUsuarios) Insert(usuario Usuario) {
	nuevo := &nodoUsuario{usuario, nil}

	if this.raiz == nil {
		this.raiz = nuevo

	} else {
		aux := this.raiz

		this.raiz = nuevo
		this.raiz.siguiente = aux
	}
	this.size++
}

func (this *ListaUsuarios) UserExist(dpi int) bool {
	aux := this.raiz

	encontro := false

	for aux != nil {
		if aux.usuario.Dpi == dpi {
			encontro = true
			break

		}
		aux = aux.siguiente
	}
	return encontro
}
func (this *ListaUsuarios) getUser(dpi int) *Usuario {
	aux := this.raiz

	for aux != nil {
		if aux.usuario.Dpi == dpi {
			return &aux.usuario

		}
		aux = aux.siguiente
	}
	return nil

}
