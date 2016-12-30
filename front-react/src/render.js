import React from 'react'
import {render} from 'react-dom'
import { Provider } from 'react-redux'
import { createStore, combineReducers } from 'redux'
import { Router, Route, browserHistory } from 'react-router'
import { syncHistoryWithStore, routerReducer } from 'react-router-redux'

import * as reducers from './reducers'

import App from './App';
import root from './engines'

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
        {root.routes}
      </Route>
    </Router>
  </Provider>,
  document.getElementById('root')
);
}

export default main
