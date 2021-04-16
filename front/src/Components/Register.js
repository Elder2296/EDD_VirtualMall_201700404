import React from 'react'

function Register() {
    return (
        <div className="centered ui middle aligned center aligned grid">
  <div className="column">
    <h2 className="ui teal image header">
      <div className="content">
        Registro 
      </div>
    </h2>
    <form className="ui large form">
      <div className="ui stacked segment">
        <div className="field">
          <div className="ui left icon input">
            <i className="user icon" />
            <input type="text" name="email" placeholder="E-mail address" />
          </div>
        </div>
        <div className="field">
          <div className="ui left icon input">
            <i className="user icon" />
            <input type="text" name="dpi" placeholder="DPI" />
          </div>
        </div>
        <div className="field">
          <div className="ui left icon input">
            <i className="user icon" />
            <input type="text" name="name" placeholder="Name" />
          </div>
        </div>
        <div className="field">
          <div className="ui left icon input">
            <i className="lock icon" />
            <input type="password" name="password" placeholder="Password" />
          </div>
        </div>
        <div className="ui fluid large teal submit button">Registrar</div>
      </div>
      <div className="ui error message" />
    </form>
    
  </div>
</div>
    )
        
}

export default Register
