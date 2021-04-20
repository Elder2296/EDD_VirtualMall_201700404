import {React, useState} from 'react'
import { useHistory } from 'react-router-dom'



const axios = require('axios').default
function Register(props) {
  let history= useHistory()
  const [email, setemail] = useState("")
  const [dpi, setdpi] = useState(0)
  const [name, setname] = useState("")
  const [pass, setpass] = useState("")

  const enviar=()=>{
    console.log("mail: "+email+" dpi: "+dpi+" name: "+name+" pass: "+pass)
    var usuario={
        Dpi:dpi,
        Nombre:name,
        Correo:email,
        Password:pass,
        Cuenta:"Usuario" 
    }

    async function postdata(){

      var user= localStorage.getItem("usuario")

      if(user==null || user==undefined){
        JSON.stringify(usuario)
        const post= await axios.post('http://localhost:3000/Registrar',usuario)
        history.push('/Login')
      }

    }
    postdata()


  }

    return (
        <div className="centered ui middle aligned center aligned grid">
  <div className="column">
    <h2 className="ui teal image header">
      <div className="content">
        Registro 
      </div>
    </h2>
    <form className="ui large form">
      <div className="ui stacked segment">
        <div className="field">
          <div className="ui left icon input">
            <i className="user icon" />
            <input type="text" name="email" placeholder="E-mail address" onChange={e=> setemail(e.target.value)} />
          </div>
        </div>
        <div className="field">
          <div className="ui left icon input">
            <i className="user icon" />
            <input type="text" name="dpi" placeholder="DPI" onChange={e=> setdpi(e.target.value)}/>
          </div>
        </div>
        <div className="field">
          <div className="ui left icon input">
            <i className="user icon" />
            <input type="text" name="name" placeholder="Name" onChange={e=> setname(e.target.value)} />
          </div>
        </div>
        <div className="field">
          <div className="ui left icon input">
            <i className="lock icon" />
            <input type="password" name="password" placeholder="Password" onChange={e=> setpass(e.target.value)} />
          </div>
        </div>
        <div className="ui fluid large teal submit button" onClick={enviar}>Registrar</div>
      </div>
      <div className="ui error message" />
    </form>
    
  </div>
</div>
    )
        
}

export default Register
