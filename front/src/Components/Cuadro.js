import {React, useState} from 'react'
import '../css/Carta.css'
import {Button} from 'semantic-ui-react'

function Cuadro(props) {
    
    const [cantidad, setcantidad] = useState(0)
    function sumar () {
        setcantidad(cantidad+1)
    }
    function restar () {
        if(cantidad==0){

        }else{
            setcantidad(cantidad-1)
        }
        
    }
    
    const enviar = () =>{
        var json={
            Id:props.tienda,
            Codigo: (props.codigo).toString(),
            Nombre:props.nombre,
            Cantidad:cantidad.toString(),
            Precio:(cantidad*props.precio).toString()


        }

        var datos = localStorage.getItem('productos')

        if(datos==null  || datos== undefined  ){
            let total=0
            total=json.Precio
            localStorage.setItem('total',JSON.stringify(total))
            localStorage.setItem('productos', JSON.stringify([json]))
            

        }else{
            datos=JSON.parse(datos)
            datos.push(json)
            let total=0
            datos.map((c,index)=>{
                total+=parseInt(c.Precio) 
            })
            console.log("el total es: "+total)
            localStorage.setItem('total',JSON.stringify(total))
            localStorage.setItem('productos', JSON.stringify(datos))
        }
        alert(JSON.stringify(json))

        console.log("ENTRO AL METOD ENVIAR")
        
    }
    return (
        <div className="column carta">
            <div className="ui card">
                <div className="image">
                    <img src={props.imagen}/>
                </div>
                <div className="content">
                    <div className="header">{props.nombre} </div>
                    
                    <div className="description">{props.descripcion}</div>
                    <div className="ui green button center fluid" onClick={enviar} >      Comprar  </div>  
                    <br></br>
                    <div className="center">
                        <Button.Group>
                            <Button  onClick={restar} >-</Button>
                            <Button.Or />
                            <Button positive onClick={sumar}>+</Button>
                            <Button color='red'>cantidad: {cantidad}</Button>
                        </Button.Group>
                    </div>
                    
           
                </div>
                <div className="extra content">
                    
                    <span className="dollar sign icon  ">Q {props.precio}</span>
                   

                </div>

            </div>
        </div>
    )
}

export default Cuadro
