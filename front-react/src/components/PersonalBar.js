import React, { PropTypes } from 'react'
import { connect } from 'react-redux'
import {MenuItem, NavDropdown} from 'react-bootstrap'
import i18n from 'i18next'
import {LinkContainer} from 'react-router-bootstrap'
import { browserHistory } from 'react-router'

import {signOut} from '../actions'
import {_delete} from '../ajax'
import {TOKEN} from '../constants'

const W = ({user, onSignOut}) => {
  if(user.uid) {
    return (<NavDropdown title={i18n.t('personal-bar.welcome', {name:user.name})} id="personal-bar">
      <LinkContainer to={{ pathname: `/dashboard` }}>
        <MenuItem>{i18n.t('personal-bar.dashboard')}</MenuItem>
      </LinkContainer>
      <MenuItem divider />
      <MenuItem onClick={onSignOut}>{i18n.t('personal-bar.sign-out')}</MenuItem>
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
  user: PropTypes.object.isRequired,
  onSignOut: PropTypes.func.isRequired
}

const M = connect(
  (state) => { return {user:state.currentUser} },
  (dispatch) => {
    return {
      onSignOut: () => {
        _delete('/users/sign-out').then( ()=>{
          dispatch(signOut())
          window.sessionStorage.removeItem(TOKEN)
          browserHistory.push('/users/sign-in')
        } )
      }
    }
  }
)(W)

export default M
