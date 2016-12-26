import React, { PropTypes } from 'react'
import { connect } from 'react-redux'

const W = ({info}) => (
  <div>
    header
    <h1>{info.title}</h1>
  </div>
)

W.propTypes = {
  info: PropTypes.object.isRequired
}

const M = connect(
  (state) => {
    return {info:state.siteInfo}
  },
  (dispatch) => { return {} }
)(W)

export default M
