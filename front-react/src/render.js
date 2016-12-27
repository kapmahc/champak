import React from 'react';
import {render} from 'react-dom';
import { Provider } from 'react-redux'
import { createStore, combineReducers } from 'redux'
import { Router, Route, IndexRoute, browserHistory } from 'react-router'
import { syncHistoryWithStore, routerReducer } from 'react-router-redux'

import './main.css';

import * as reducers from './reducers'

import App from './App';
import {SignIn, SignUp, Confirm, Unlock, ForgotPassword, ResetPassword} from './auth/users'
import Home from './components/Home'

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
        <Route path="users/sign-in" component={SignIn}/>
        <Route path="users/sign-up" component={SignUp}/>
        <Route path="users/confirm" component={Confirm}/>
        <Route path="users/unlock" component={Unlock}/>
        <Route path="users/forgot-password" component={ForgotPassword}/>
        <Route path="users/reset-password/:token" component={ResetPassword}/>
      </Route>
    </Router>
  </Provider>,
  document.getElementById('root')
);
}

export default main
