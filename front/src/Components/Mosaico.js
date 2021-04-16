import React from 'react'
import Carta from './Carta'
import '../css/Fondo.css'

function Mosaico(props) {
    //console.log(props)
    return (
        <>
        

        
        <div className=" ui segment  container  grid">
            <div className="ui four column link cards row">
            
                {
                    props.tiendas.map((c,index) => (
                        <Carta nombre={c.Nombre}
                            logo={c.Logo}
                            descripcion={c.Descripcion}
                            id={c.Id}
                            contacto={c.Contacto}
                            calificacion={c.Calificacion}
                            key={c.Id}
                        />
                    ))
                }

            </div>
            
        </div>
        </>
    )
}

export default Mosaico