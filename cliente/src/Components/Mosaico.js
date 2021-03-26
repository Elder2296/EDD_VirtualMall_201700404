import React from 'react'
import Carta from './Carta'

function Mosaico(props) {
    return (
        <div className="ui segment mosaico container">
            <div className="ui three column link cards row">
                {
                    props.tiendas.map((c,index) => (
                        <Carta nombre={c.nombre}
                            logo={c.Logo}
                            descripcion={c.Descripcion}
                            id={c.id}
                            contacto={c.Contacto}
                            key={c.id}
                        />
                    ))
                }

            </div>
            
        </div>
    )
}

export default Mosaico
