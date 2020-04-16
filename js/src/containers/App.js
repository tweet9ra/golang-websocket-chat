import { connect } from 'react-redux'
import ChatContainer from "./ChatContainer";
import Login from "./Login";
import React from "react";

const App = ({isAuthorized}) => {
  return (
    <div className="App">
      <header>
        <h1>GO + React WebSocket chat</h1>
      </header>
      {
        isAuthorized
          ? <ChatContainer/>
          : <Login/>
      }
      <footer>
        <h3>Design, HTML & CSS template credits <a
          href="https://codepen.io/drehimself/pen/KdXwxR">https://codepen.io/drehimself/pen/KdXwxR</a></h3>
      </footer>
    </div>
  )
}

const mapStateToProps = state => {
  return {
    isAuthorized: !!state.auth.user
  }
}


export default connect(mapStateToProps)(App)