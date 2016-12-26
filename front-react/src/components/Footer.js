import React, { PropTypes } from 'react'
import { connect } from 'react-redux'
import i18n from 'i18next'

import {LOCALE} from '../constants'

const W = ({info}) => (
  <div className="row">
    <hr/>
    <footer>
      <p className="pull-right">
        {i18n.t('footer.other-languages')}
        {info.languages.map( (l, i) => (
          <a className="block" href={`/?${LOCALE}=${l}`} key={i}>{i18n.t(`languages.${l}`)}</a>
        ))}
      </p>
      <p>
        &copy; {info.copyright} &nbsp;
        <a href="#">Privacy</a> &middot;
        <a href="#">Terms</a>
      </p>
    </footer>
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
