import auth from './auth'
import site from './site'

export default{
  routes: [].concat(auth.routes)
    .concat(site.routes)
}
