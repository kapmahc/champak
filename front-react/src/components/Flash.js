import React, {PropTypes} from 'react'
import { connect } from 'react-redux'
import {Alert} from 'react-bootstrap'
import i18n from 'i18next'
import {hideFlash} from '../actions'

const W = ({onHide, msg}) => {
  var sty
  switch (msg.type) {
    case 'notice':
      sty = 'success'
      break;
    case 'alert':
      sty = 'danger'
      break
    default:
      sty = 'info'
  }
  if(msg.show){
    return (<div className="row">
      <br/>
      <Alert bsStyle={sty} onDismiss={onHide}>
        <h4>{i18n.t(`flashs.${msg.type}`)}</h4>
        <p>{msg.body}</p>
      </Alert>
    </div>)
  }
  return <div className="row"/>
}

W.propTypes = {
  msg: PropTypes.object.isRequired,
  onHide: PropTypes.func.isRequired
}

const M = connect(
  (state) => {
    return {msg: state.flash}
  },
  (dispatch) => {
    return {
      onHide: () => {
        dispatch(hideFlash())
      },
    }
  }
)(W)

export default M
