import React from 'react'
import {Route} from 'react-router'

import {SignIn, SignUp, Confirm, Unlock, ForgotPassword, ResetPassword} from './non-sign-in'
import {Info, ChangePassword, Logs} from './profiles'
import Dashboard from '../../Dashboard'

const W = [
  <Route key="auth.users.non-sign-in" path="users">
    <Route path="sign-in" component={SignIn}/>
    <Route path="sign-up" component={SignUp}/>
    <Route path="confirm" component={Confirm}/>
    <Route path="unlock" component={Unlock}/>
    <Route path="forgot-password" component={ForgotPassword}/>
    <Route path="reset-password/:token" component={ResetPassword}/>
  </Route>,
  <Route key="auth.users.profiles" path="users" component={Dashboard}>
    <Route path="info" component={Info}/>
    <Route path="change-password" component={ChangePassword}/>
    <Route path="logs" component={Logs}/>
  </Route>
]

export default W
