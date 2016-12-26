import React, { PropTypes } from 'react'
import { connect } from 'react-redux'
import {MenuItem, NavDropdown} from 'react-bootstrap'
import i18n from 'i18next'

import {LOCALE} from '../constants'

const W = ({info}) => (
  <NavDropdown title={i18n.t('language-bar.switch')} id="language-bar">
    {info.languages.map( (l, i) => (
      <MenuItem key={i} href={`/?${LOCALE}=${l}`}>
        {i18n.t(`languages.${l}`)}
      </MenuItem>
    ))}
  </NavDropdown>
)

W.propTypes = {
  info: PropTypes.object.isRequired
}

const M = connect(
  (state) => { return {info:state.siteInfo} },
  (dispatch) => { return {} }
)(W)

export default M
