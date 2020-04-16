import React from "react"
import {connect} from 'react-redux'
import {authenticate} from "../actions";

const Login = ({dispatch}) => {
  let email
  let password

  return (
    <div className='container login'>
      <form onSubmit={e => {
        e.preventDefault()
        if (!email.value.trim() || !password.value.trim()) {
          return
        }
        dispatch(authenticate(email.value, password.value))
      }}>
        <div className="form-row">
          <input type="email" defaultValue="tweet9ra@mail.ru" required ref={node => email = node}/>
          <span>Email</span>
        </div>
        <div className="form-row">
          <input type="password" required ref={node => password = node}/>
          <span>Password</span>
        </div>
        <div className="form-row"></div>
        <div className="form-row">
          <button type="submit">Login to your Account!</button>
        </div>
      </form>
    </div>
  )
}

export default connect()(Login)