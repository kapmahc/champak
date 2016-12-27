import React from 'react'
import {Route, IndexRoute} from 'react-router'

import Home from './Home'
import NoMatch from './NoMatch'

const W = [
  <IndexRoute key="home" component={Home}/>,  
  <Route key="no-match" path="*" component={NoMatch}/>
]

export default W
