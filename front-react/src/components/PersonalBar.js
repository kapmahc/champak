import React, { PropTypes } from 'react'
import { connect } from 'react-redux'
import {MenuItem, NavDropdown} from 'react-bootstrap'
import i18n from 'i18next'
import {LinkContainer} from 'react-router-bootstrap'

const W = ({user}) => {
  if(user.uid) {
    // TODO
    return (<NavDropdown title={i18n.t('language-bar.switch')} id="personal-bar">
      {['dashboard','sign-out'].map( (l, i) => (
        <MenuItem key={i} href='/'>
          {i18n.t(`languages.${l}`)}
        </MenuItem>
      ))}
    </NavDropdown>)
  }else{
    return (<NavDropdown title={i18n.t('personal-bar.sign-in-or-up')} id="personal-bar">
      {['sign-in', 'sign-up', 'forgot-password', 'confirm', 'unlock'].map( (v, i) => (
        <LinkContainer key={i} to={{ pathname: `/users/${v}` }}>
          <MenuItem>{i18n.t(`auth.users.${v}.title`)}</MenuItem>
        </LinkContainer>
      ))}
    </NavDropdown>)
  }
}

W.propTypes = {
  user: PropTypes.object.isRequired
}

const M = connect(
  (state) => { return {user:state.currentUser} },
  (dispatch) => { return {} }
)(W)

export default M
