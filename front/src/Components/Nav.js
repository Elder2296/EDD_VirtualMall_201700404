import React, { Component } from 'react'
import { Menu, Segment } from 'semantic-ui-react'
import {BrowserRouter as Router, Switch,Route, Link} from "react-router-dom"

export default class MenuExampleInvertedSecondary extends Component {
  state = { activeItem: 'home' }

  handleItemClick = (e, { name }) => this.setState({ activeItem: name })

  render() {
    const { activeItem } = this.state

    return (
        <>
        <Segment inverted>
          <Menu inverted pointing secondary>
            <Menu.Item
              name='Home'
              active={activeItem === 'Home'}
              onClick={this.handleItemClick}
              as={Link} to={'/'}
            />
            <Menu.Item
              name='Tiendas'
              active={activeItem === 'Tiendas'}
              onClick={this.handleItemClick}
              as={Link} to={'/Tiendas'}
            />
            <Menu.Item
              name='Vision'
              active={activeItem === 'Vision'}
              onClick={this.handleItemClick}
              as={Link} to={'/Vision'}
            />
            <Menu.Item
              name='Reportes'
              active={activeItem === 'Reportes'}
              onClick={this.handleItemClick}
              as={Link} to={'/Reportes'}
            />

            <div class="right item">
              <Link to={'/Login'}><a class="ui inverted button">Log in</a></Link>
            
              <Link to={'/Registro'}><a class="ui inverted button">Sign Up</a></Link>
              <Link to={`/Carrito`}> <a class="ui inverted button"><i  class="mind dolly icon own-class">Carrito</i></a></Link>
            </div>
          </Menu>
          
        </Segment>
        
        </>
    )
  }
}

/*
<Menu inverted className="Nav">
            {
                colores.map((c,index)=>(

                    <Menu.Item as={Link} to ={url[index]}
                        key={c}
                        name={opciones[index]}
                        active={activo===c}
                        color={c}
                        onClick={()=>setactivo(c)}

                    
                    />

                ))
            }

        <Menu.Menu position='right'>
        <Button.Group>
            
        <Link to={`/Carrito`}> <Button color='purple' ><i  class="big dolly icon own-class"></i></Button></Link>
           
                           
        </Button.Group>
        
          
          
        </Menu.Menu>
        </Menu>
*/ 