import { w3cwebsocket as W3CWebSocket } from "websocket";
import {store} from "../index";
import {wsNewMessage} from "../actions";

let client

export const wsConnect = token => {
  if (client) {
    return client
  }
  client = new W3CWebSocket('ws://localhost:8000/ws')

  client.onopen = () => {
    sendMessage('AUTH', {token})
  };

  client.onmessage = message => store.dispatch(wsNewMessage(JSON.parse(message.data)))

  return client
}

export const wsDisconnect = () => {
  client.close()
  client = null
}

export const sendMessage = (actionType, data) => {
  client.send(JSON.stringify({action: actionType, data}))
}