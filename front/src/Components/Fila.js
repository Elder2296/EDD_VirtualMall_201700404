import React from 'react'

function Fila(props) {
    return (
        <>
        <tr>
            <td>{props.nombre}</td>
            <td>{props.cantidad}</td>
            <td>{props.subtotal}</td>
            <td><button class="ui red button">Eliminar</button></td>
            
        </tr>

        </>
    )
}

export default Fila
