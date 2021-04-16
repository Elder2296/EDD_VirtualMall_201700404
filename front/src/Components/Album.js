import React from 'react'
import Cuadro from './Cuadro'

function Album(props) {
    
    return (
        <>
        <br/>
        <br/>
        <br/>
        <br/>
        <br/><br/>
        <div className=" ui segment  container  grid">
            <div className="ui three column link cards row">
                {

                  props.productos.map((c,index)=>(
                      <Cuadro
                            nombre={c.Nombre}
                            descripcion={c.Descripcion}
                            precio={c.Precio}
                            imagen={c.Imagen}
                            codigo={c.Codigo}
                            tienda={props.Tienda}
                            key={c.Codigo}
                      />
                  ))  

                }

            </div>
            
        </div>
        </>
    )
}

export default Album
