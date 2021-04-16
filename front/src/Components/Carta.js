import React from 'react'
import '../css/Carta.css'
import {BrowserRouter as Router, Switch,Route, Link} from "react-router-dom"
import { Redirect} from "react-router-dom";
import Producto from "./Producto"



const axios =require('axios').default
function Carta(props) {
    const enviar=()=>{
       
        /*
        
        var infotienda={
            Calificacion:props.calificacion,
            Id:props.id
        }
        JSON.stringify([infotienda])
        console.log(infotienda)
        axios.post('http://localhost:3000/Verproductos',infotienda)
            *
        

        
        */
    }

    return (
        <>
        <div className="column carta">
            <div className="ui  card">
                <div className="image">
                    <img src={props.logo}/>
                </div>
                <div className="content">
                    <div className="header">{props.nombre} </div>
                    
                    <div className="description">{props.descripcion}</div>
                    <Link to={`/Tienda/${props.id}`}> <div className="ui olive button center fluid"  >Ver Productos</div></Link>               
                </div>
                <div className="extra content">
                    
                    <span className="dollar sign icon ">Contacto {props.contacto}</span>

                </div>

            </div>
        </div>
        <br></br>
               


        </>
    )
}

export default Carta