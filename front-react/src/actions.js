import {SIGN_IN, SIGN_OUT, REFRESH, FLASH} from './constants'

export const hideFlash = () => {
  return {
    type: FLASH,
    body: {
      type: null,
      show: false,
      body: null
    }
  }
}

export const noticeFlash = (body) => {
  return {
    type: FLASH,
    body: {
      type: 'notice',
      show: true,
      body
    }
  }
}

export const alertFlash = (body) => {
  return {
    type: FLASH,
    body: {
      type: 'alert',
      show: true,
      body
    }
  }
}


export const refresh = (info) => {
  return {
    type: REFRESH,
    info
  }
}

export const signIn = (token) => {
  return {
    type: SIGN_IN,
    token
  }
}

export const signOut = () => {
  return {
    type: SIGN_OUT
  }
}
