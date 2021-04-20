import {React, useState  } from 'react'
import { useHistory } from 'react-router-dom'
import '../css/login.css'
import {BrowserRouter as Router, Switch,Route, Link, Redirect} from "react-router-dom"

const axios = require('axios').default

function Login() {
    let history = useHistory()
    const [dpi, setdpi] = useState(0)
    const [Contraseña, setContraseña] = useState("")

    const enviar =()=>{
      console.log("Dpi: "+dpi+" Password: "+Contraseña)
      var datos={
        Dpi:dpi,
        Password:Contraseña
      }
      JSON.stringify(datos)
      var user=localStorage.getItem("usuario")
      async function getTiendas(){
        console.log("Entro aca 1")
          if(user ==null || user==undefined){
            console.log("Entro aca 2")
              const info = await axios.post('http://localhost:3000/Validar',datos)

              console.log("RETORNO: ", info)
              console.log(info.data.Cuenta)
              var informacion={
                Dpi:info.data.Dpi,
                Usuario:info.data.Cuenta
              }
              if(info.data.Cuenta==="Usuario"){
                
                localStorage.setItem("Usuario",JSON.stringify(informacion))
                console.log("Entro al usuario")
                history.push('/Tiendas')
                
              }else if(info.data.Cuenta=="Admin"){
                localStorage.setItem("Usuario",JSON.stringify(informacion))
                console.log("Entro al admin")
                
                history.push('/Reportes')

              }else{
                console.log("Entro por este lado")
              }
              
          }
      }
      getTiendas()
      
    }


    return (
        <div class="centered ui middle aligned center aligned grid">
  <div class="column">
    <h2 class="ui teal image header">
      
      <div class="content">
        Log-in 
      </div>
    </h2>
    <form class="ui large form">
      <div class="ui stacked segment">
        <div class="field">
          <div class="ui left icon input">
            <i class="user icon"></i>
            <input type="text" name="email" placeholder="Dpi" onChange={e=> setdpi(e.target.value)} />
          </div>
        </div>
        <div class="field">
          <div class="ui left icon input">
            <i class="lock icon"></i>
            <input type="password" name="password" placeholder="Password" onChange={e=> setContraseña(e.target.value)}  />
          </div>
        </div>
        <div class="ui fluid large teal submit button" onClick={enviar}>Login</div>
      </div>

      <div class="ui error message"></div>

    </form>

    <div class="ui message">
      New to us? <Link to={'/Registro'}><a href="#">Sign Up</a></Link>
    </div>
  </div>
</div>

    )
}

export default Login
