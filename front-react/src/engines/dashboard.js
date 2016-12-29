import auth from './auth/links'
import site from './site/links'

const items = (user) => {
  return [].concat(auth(user))
    .concat(site(user))
}

export default items
