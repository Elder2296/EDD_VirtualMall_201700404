import React from 'react'

function Carta(props) {
    return (
        <div className="column carta">
            <div className="ui card">
                <div className="image">
                    <img src={props.logo}/>
                </div>
                <div className="content">
                    <div className="header">{props.nombre} </div>
                    
                    <div className="description">{props.descripcion}</div>
                    <div className="ui basic green button center fluid" onClick={() => {console.log(props.id)}} > Comprar</div>              
                </div>
                <div className="extra content">
                    
                    <span className="dollar sign icon ">Contacto {props.contacto}</span>

                </div>

            </div>
        </div>
    )
}

export default Carta
