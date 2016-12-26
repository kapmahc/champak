export const currentUser = (state={}, action) => {
  switch (action.type) {
    case 'SIGN_IN':
      return action.user
    case 'SIGN_OUT':
      return {}
    default:
      return state
  }
}

export const siteInfo = (state={title:'', copyright:''}, action) => {
  switch(action.type) {
    case 'REFRESH':
      return action.info
    default:
      return state
  }
}
