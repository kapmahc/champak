import React, { PropTypes } from 'react'
import { connect } from 'react-redux'

const SignInW = ({info}) => (
  <div>
    <h1>{info.copyright}</h1>
    sign in
  </div>
)

SignInW.propTypes = {
  info: PropTypes.object.isRequired
}

export const SignIn = connect(
  (state) => { return {info:state.siteInfo} },
  (dispatch) => { return {} }
)(SignInW)

// -----------------------------------------------------------------------------

const SignUpW = ({info}) => (
  <div>
    <h1>{info.copyright}</h1>
    sign up
  </div>
)

SignUpW.propTypes = {
  info: PropTypes.object.isRequired
}

export const SignUp = connect(
  (state) => { return {info:state.siteInfo} },
  (dispatch) => { return {} }
)(SignUpW)
