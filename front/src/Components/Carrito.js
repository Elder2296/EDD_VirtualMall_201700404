import {React,useState} from 'react'
import Tabla from './Tabla'
import '../css/Tiendas.css'

function Carrito() {
    
    var datos =localStorage.getItem('productos')

    if(datos==null || datos == undefined){

    }else{
        datos=JSON.parse(datos)
        //console.log("Contenido del carrito")
        //console.log(datos)
        

        
    }
    
    

       
        

    
        return (
        
            <Tabla datos={datos} ></Tabla>
        
        
        )
    

    
}

export default Carrito
