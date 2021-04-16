import {React, useState, useEffect} from 'react'
import '../css/Tiendas.css'
import '../css/Carga.css'
import Mosaico from './Mosaico'

const axios =require('axios').default

function Tiendas() {
    const [tiendas, setTiendas] = useState([])
    const [loading,setloading] = useState(false)

    useEffect(() => {
        async function getTiendas(){
            if(tiendas.length===0){
                const data= await axios.get('http://localhost:3000/Tiendas')
                //console.log(data)
                //console.log("Estos ocurrio")
                setTiendas(data.data)
            
                
                setloading(true)
            }
        }
        getTiendas()
    }
    
    );
    
    if(loading ===false){
        return (
            <div className="ui segment carga">
                <div className="ui active dimmer ">
                    <div className="ui text loader">

                    </div>

                </div>

            </div>

        )

    }else{
        return (
            <div className="tiendas">
                
                <br></br>
                <br></br>
                <br></br>
                <br></br>
                <br></br>
                
                
                <Mosaico tiendas={tiendas}></Mosaico>
                
            </div>
        )    

    }



    
}

export default Tiendas
