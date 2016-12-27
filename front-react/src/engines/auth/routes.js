import React from 'react'
import {Route} from 'react-router'

import {SignIn, SignUp, Confirm, Unlock, ForgotPassword, ResetPassword} from './non-sign-in'

const W = [
  <Route key="auth.users" path="users">
    <Route path="sign-in" component={SignIn}/>
    <Route path="sign-up" component={SignUp}/>
    <Route path="confirm" component={Confirm}/>
    <Route path="unlock" component={Unlock}/>
    <Route path="forgot-password" component={ForgotPassword}/>
    <Route path="reset-password/:token" component={ResetPassword}/>
  </Route>
]

export default W
