import React from 'react'
import {render} from 'react-dom'
import { Provider } from 'react-redux'
import { createStore, combineReducers } from 'redux'
import { Router, Route, IndexRoute, browserHistory } from 'react-router'
import { syncHistoryWithStore, routerReducer } from 'react-router-redux'

import './main.css';

import * as reducers from './reducers'

import App from './App';
import Auth from './auth/routes'
import Home from './components/Home'
import NoMatch from './components/NoMatch'

let store = createStore(
  combineReducers({
    ...reducers,
    routing: routerReducer
  })
)
let history = syncHistoryWithStore(browserHistory, store)

const main = () => {
render(
  <Provider store={store}>
    <Router history={history}>
      <Route path="/" component={App}>
        <IndexRoute component={Home}/>
        {Auth}
        <Route path="*" component={NoMatch}/>
      </Route>
    </Router>
  </Provider>,
  document.getElementById('root')
);
}

export default main
