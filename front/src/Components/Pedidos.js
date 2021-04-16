import {React, useEffect} from 'react'

const axios = require('axios').default

function Pedidos() {
    
    useEffect( () =>{
        var datos= localStorage.getItem('productos')
        
            
            if(datos==null || datos==undefined){

            }else{
                async function hacerpedido(){
                
                var info={
                    Total:localStorage.getItem('total'),
                    Datos:localStorage.getItem('productos')
                    
                }
                JSON.stringify(info)
                console.log("Lo que se envia al server")
                console.log(info)
                await axios.post('http://localhost:3000/HacerPedido',info)

                localStorage.clear();

                
                
            }
            hacerpedido()
        }
        
    })


    return (
        <div>
            Estas en Pedidos
        </div>
    )
}

export default Pedidos
