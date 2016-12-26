export const REFRESH = 'REFRESH'
export const SIGN_IN = 'SIGN_IN'
export const SIGN_OUT = 'SIGN_OUT'

export const refresh = (info) => {
  return {
    type: 'REFRESH',
    info
  }
}

export const signIn = (user) => {
  return {
    type: 'SIGN_IN',
    user
  }
}

export const signOut = () => {
  return {
    type: 'SIGN_OUT'
  }
}
