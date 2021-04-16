import {React} from 'react'

import Fila from './Fila'
import {BrowserRouter as Router, Switch,Route, Link} from "react-router-dom"

import {Button} from 'semantic-ui-react'

const axios= require('axios').default

function Tabla(props) {
    /*function enviarDatos(){
        console.log("entro al evento")
        var data={
            Datos:localStorage.getItem('productos'),
            Total:localStorage.getItem('total')
        }
        async function post(){
            JSON.stringify(data)
            const info = await axios.post('http://localhost:3000/HacerPedido',data
                )

            console.log(info)

        }
        
    }*/

    let monto=0
    var datos =localStorage.getItem('total')

    if(datos==null || datos == undefined){

    }else{
        datos=JSON.parse(datos)
        console.log(datos)
        monto=datos

        
    }
    

    return (
        <>
        <table className="ui inverted teal table">
            <thead>
                <tr>
                    <th>Producto</th>
                    <th>Cantidad</th>
                    <th>Costo</th>
                    <th></th>
                </tr>
            </thead>
            <tbody>
            {
                props.datos.map((c,index)=>(
                    <Fila 
                        idproducto={c.Codigo}
                        nombre={c.Nombre}
                        cantidad={c.Cantidad}
                        subtotal={c.Precio}
                        
                    />

                ))
            }

                <tr>
                    <th>Total</th>
                    <th></th>
                    <th>Q {parseFloat(monto) }</th>
                    <th></th>
                </tr>
            </tbody>

        </table>
        <div className="container">
        <div class="ui buttons center">
            <Link to={`/Tiendas`}><button class="ui button">Seguir Comprando</button></Link>
            
            <div class="or"></div>
            <Link to={`/Pedidos`}><button class="ui positive button" >Generar Pedido</button> </Link>
        </div>

        </div>
        </>
    )
}

export default Tabla
