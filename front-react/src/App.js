import React, {PropTypes, Component} from 'react'
import { connect } from 'react-redux'

import Header from './components/Header'
import Footer from './components/Footer'
import Flash from './components/Flash'
import {refresh, signIn, alertFlash, noticeFlash} from './actions'
import {get} from './ajax'
import {TOKEN} from './constants'

class W extends Component {
  componentDidMount() {
    const {onRefresh} = this.props
    onRefresh(this.props.location.query)
  }
  render () {
    const {children} = this.props
    return <div>
      <Header />
      <div className="container">
        <Flash />
        {children}
        <Footer />
      </div>
    </div>
  }
}

W.propTypes = {
  onRefresh: PropTypes.func.isRequired,
  children: PropTypes.node.isRequired
}

const M = connect(
  (state) => {
    return {}
  },
  (dispatch) => {
    return {
      onRefresh: (q) => {
        var tkn = window.sessionStorage.getItem(TOKEN)
        if(tkn){
          dispatch(signIn(tkn))
        }
        // ---------
        get('/site/info', ).then(
          (rst)=> { dispatch(refresh(rst)) }
        )
        // --------
        if(q.alert){
          dispatch(alertFlash(q.alert))
        }
        if(q.notice){
          dispatch(noticeFlash(q.notice))
        }
      }
    }
  }
)(W)

export default M
