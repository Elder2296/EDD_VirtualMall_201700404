import React from 'react'

function Reportes() {

    

    var user=localStorage.getItem('Usuario')

    if(user==null || user==  undefined){
        return(
            <div>
               <h1>No se ha iniciado Sesion</h1> 
            </div>
        )

    }else{
        user = JSON.parse(user)

        if(user.Usuario=='Admin'){
            return (
                <div>
                    Aqui van los reportes
                </div>
            )
        }else{
            return(
                <div>
                   <h1>No tienes acceso</h1> 
                </div>
                )
        }
    }
    
}

export default Reportes
