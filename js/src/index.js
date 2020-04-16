import React from 'react'
import ReactDOM from 'react-dom'
import App from './containers/App'
import * as serviceWorker from './serviceWorker'
import rootReducer from './reducers'
import {Provider} from 'react-redux'
import {applyMiddleware, createStore} from 'redux'
import './scss/chat.scss'
import './scss/app.scss'
import './scss/login.scss'
import thunkMiddleware from "redux-thunk";
import {wsConnect} from "./helpers/wsConnection";

const currentUser = JSON.parse(localStorage.getItem('user'))
const state = {};
if (currentUser) {
  state.auth = {
    user: currentUser
  }
  wsConnect(currentUser.token)
}
state.chats = null

export const store = createStore(rootReducer, state, applyMiddleware(thunkMiddleware))


ReactDOM.render(
  <Provider store={store}>
    <React.StrictMode>
      <App/>
    </React.StrictMode>
  </Provider>,
  document.getElementById('root')
);

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.unregister();
