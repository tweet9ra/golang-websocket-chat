import {fetchGet, fetchPost} from "../helpers/fetch";
import {wsConnect, wsDisconnect} from "../helpers/wsConnection";

const JSAlert = require("js-alert");

export const setUser = user => {
  localStorage.setItem('user', JSON.stringify(user))
  return {
    type: 'SET_USER',
    user: user
  }
}

export const logout = () => dispatch => {
  wsDisconnect()
  dispatch(setUser(null))
}

export const authenticate = (email, password) => (dispatch, getState) => {
  fetchPost('/user/login', {email, password})
    .then(data => {
      if (data.status) {
        dispatch(setUser(data.data))
        wsConnect(data.data.token)
      } else {
        JSAlert.alert(data.message)
      }
    })
}

const setChats = chats => ({
  type: 'SET_CHATS',
  chats
})

export const loadChatsIfNotLoaded = () => (dispatch, getState) => {
  if (getState().chats === null) {
    fetchGet('/chats')
      .then(data => {
        if (!data.status) {
          JSAlert.alert(data.message)
        }

        dispatch(setChats(data.chats))
      })
  }
}

export const wsNewMessage = message => dispatch => {
  switch (message.action) {
    case 'AUTH':
      if (!message.data.result) {
        JSAlert.alert("WebSocket authentication failed: " + message.data.error)
      }
      break;
    case 'NEW_MESSAGE':
      dispatch({
        type: 'NEW_MESSAGE',
        message: message.data
      })
      break;
    default:
      break;
  }
}