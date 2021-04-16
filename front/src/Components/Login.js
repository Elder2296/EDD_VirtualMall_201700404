import React from 'react'
import '../css/login.css'
import {BrowserRouter as Router, Switch,Route, Link} from "react-router-dom"

function Login() {
    return (
        <div class="centered ui middle aligned center aligned grid">
  <div class="column">
    <h2 class="ui teal image header">
      
      <div class="content">
        Log-in 
      </div>
    </h2>
    <form class="ui large form">
      <div class="ui stacked segment">
        <div class="field">
          <div class="ui left icon input">
            <i class="user icon"></i>
            <input type="text" name="email" placeholder="E-mail address"/>
          </div>
        </div>
        <div class="field">
          <div class="ui left icon input">
            <i class="lock icon"></i>
            <input type="password" name="password" placeholder="Password"/>
          </div>
        </div>
        <div class="ui fluid large teal submit button">Login</div>
      </div>

      <div class="ui error message"></div>

    </form>

    <div class="ui message">
      New to us? <Link to={'/Registro'}><a href="#">Sign Up</a></Link>
    </div>
  </div>
</div>

    )
}

export default Login
