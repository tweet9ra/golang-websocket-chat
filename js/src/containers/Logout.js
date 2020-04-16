import {logout} from "../actions";
import React from "react";
import {connect} from 'react-redux'

const Logout = ({logout}) => <div className="logout" onClick={logout}><button>Logout</button></div>

const mapStateToProps = state => ({
})

const mapDispatchToProps = dispatch => ({
  logout: () => dispatch(logout())
})

export default connect(mapStateToProps, mapDispatchToProps)(Logout)