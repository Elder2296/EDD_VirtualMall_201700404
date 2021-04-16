import {React, useState, useEffect} from 'react'
import { useParams } from 'react-router-dom'
import '../css/Carga.css'
import '../css/Tiendas.css'
import Albun from './Album'



const axios = require('axios').default
function Producto( ) {
    const {id} = useParams();
    const [productos, setproductos] = useState([])
    const [loading, setloading] = useState(false)

    useEffect( () => {
        async function getProductos(){
            if(productos.length===0){
                var info ={
                    Id:id
                }
                JSON.stringify(info)
                const data = await axios.post('http://localhost:3000/Verproductos',info)
                console.log(data)
                setproductos(data.data)
                /*.then(res=>{
                    
                    
                    console.log(res);
                     
                })
                .catch(error=>{
                    console.log(error)
                })*/
                console.log(info)
                setloading(true)
                

            }
        }
        getProductos()

    })
    if(loading===false){
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
                <Albun productos={productos} Tienda={id}></Albun>
            </div>
        )

    }
    
}

export default Producto
