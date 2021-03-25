import React from 'react'
import {BrowserRouter as Router,Route} from 'react-router-dom'
import Navbar from './Components/Nav'
import Home from './Components/Home'
import Tiendas from './Components/Tiendas'

function App() {
  return (
    <Router>
      <Navbar/>
      <Route path="/"  component={Home}/>
      <Route path="/Tiendas" component={Tiendas} />
      <Route path="/Pedidos"  />
      

      

    </Router>
  )
}

export default App
