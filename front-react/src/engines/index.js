import auth from './auth'
import site from './site'

export default{
  dashboard(user){
    return [].concat(auth.dashboard(user))
      .concat(site.dashboard(user))
  },
  routes: [].concat(auth.routes)
    .concat(site.routes)
}
