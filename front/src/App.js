import React from 'react'
import {BrowserRouter as Router,Route} from 'react-router-dom'
import Navbar from './Components/Nav'
import Home from './Components/Home'
import Producto from './Components/Producto'
import Tiendas from './Components/Tiendas'
import Carrito from './Components/Carrito'
import Pedidos from './Components/Pedidos'
import Login from './Components/Login'
import Registro from './Components/Register'
function App() {
  localStorage.clear();
  return (
    <>
      
      <Router>

      
      <div class="pusher">
        <div class="ui inverted vertical masthead center aligned segment">

        <Navbar/>

        <Route exact path="/"  component={Home}/>
        <Route exact path="/Tiendas" component={Tiendas} />
        <Route  path="/Tienda/:id"  >  <Producto></Producto>  </Route>
        <Route  path="/Carrito"  ><Carrito></Carrito></Route>
        <Route  path="/Pedidos"  ><Pedidos/></Route>
        <Route  path="/Login"><Login></Login></Route>
        <Route  path="/Registro"><Registro></Registro></Route>

      







        
        
    </div>


  
              
    </div>

    </Router>

    </>
    
  )
}

export default App

/*<Router>
      <Navbar/>
      <Route exact path="/"  component={Home}/>
      <Route exact path="/Tiendas" component={Tiendas} />
      <Route  path="/Tienda/:id"  >  <Producto></Producto>  </Route>
      <Route  path="/Carrito"  ><Carrito></Carrito></Route>
      <Route  path="/Pedidos"  ><Pedidos/></Route>
    </Router>*/