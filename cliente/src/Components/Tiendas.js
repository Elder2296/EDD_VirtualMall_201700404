import {React, useState, useEffect} from 'react'

import '../css/Tiendas.css'
import Mosaico from './Mosaico'

const axios =require('axios').default

function Tiendas() {
    const [tiendas, setTiendas] = useState([])
    const [loading,setloading] = useState(false)

    useEffect(() => {
        async function getTiendas(){
            if(tiendas.length===0){
                const data= await axios.get('http://localhost:3000/Tiendas')
                console.log()
                setTiendas(data)
                setloading(true)
            }
        }
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
                <Mosaico tiendas={tiendas}></Mosaico>
                
            </div>
        )    

    }



    
}

export default Tiendas
