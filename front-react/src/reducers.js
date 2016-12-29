import jwtDecode from 'jwt-decode'

import {SIGN_IN, SIGN_OUT, REFRESH, FLASH} from './constants'

export const currentUser = (state={}, action) => {
  switch (action.type) {
    case SIGN_IN:
      var user = jwtDecode(action.token)
      user.is_sign_in = function(){return this.uid != null}
      return user
    case SIGN_OUT:
      return {}
    default:
      return state
  }
}

export const siteInfo = (state={languages: [], top:[], bottom: []}, action) => {
  switch(action.type) {
    case REFRESH:
      return action.info
    default:
      return state
  }
}

export const flash = (state={}, action) => {
  switch (action.type) {
    case FLASH:
      return action.body
    default:
      return state
  }
}
