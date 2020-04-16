import fetch from 'cross-fetch'
import {store} from '../index'

const apiEndpoint = 'http://localhost:8000/api'


export function fetchPost(urlPath, body) {
  const params = {
    method: 'POST',
    body: JSON.stringify(body)
  }

  addAuthorizationHeader(params)

  return fetch(apiEndpoint + urlPath, params).then(response => response.json())
}

export function fetchGet(urlPath) {
  const params = {
    method: 'GET'
  }

  addAuthorizationHeader(params)

  return fetch(apiEndpoint + urlPath, params).then(response => response.json())
}

const addAuthorizationHeader = (params) => {
  const state = store.getState()
  if (state.auth.user && state.auth.user.token) {
    params.headers = {'Authorization': `Bearer ${state.auth.user.token}`}
  }

}