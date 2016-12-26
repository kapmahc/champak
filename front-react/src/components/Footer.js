import React, { PropTypes } from 'react'
import { connect } from 'react-redux'

const W = ({info}) => (
  <div>
    <h1>{info.copyright}</h1>
    footer
  </div>
)

W.propTypes = {
  info: PropTypes.object.isRequired
}

const M = connect(
  (state) => { return {info:state.siteInfo} },
  (dispatch) => { return {} }
)(W)

export default M
