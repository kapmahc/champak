import React from 'react'
import {Route, IndexRoute} from 'react-router'

import Home from './Home'
import Dashboard from './Dashboard'
import NoMatch from './NoMatch'

const W = [
  <IndexRoute key="index" component={Home}/>,
  <Route key="home" path="home" component={Home}/>,
  <Route key="dashboard" path="dashboard" component={Dashboard}/>,  
  <Route key="no-match" path="*" component={NoMatch}/>
]

export default W
