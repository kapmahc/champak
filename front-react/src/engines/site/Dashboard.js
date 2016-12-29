import React, { PropTypes } from 'react'
import { connect } from 'react-redux'
import i18n from 'i18next'

import dashboard from '../dashboard'

const W = React.createClass({
  render(){
    const {user} = this.props
    var links = dashboard(user)
    console.log(links)    
    return (<div className="row">
      <h2>{i18n.t('header.dashboard')}</h2>
      <hr/>
    </div>)
  }
})


W.propTypes = {
  user: PropTypes.object.isRequired
}

const M = connect(
  (state) => {
    return {
      user: state.currentUser
    }
  },
  (dispatch) => { return {} }
)(W)

export default M
