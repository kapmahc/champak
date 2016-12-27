import React, {PropTypes, Component} from 'react'
import { connect } from 'react-redux'

import Header from './components/Header'
import Footer from './components/Footer'
import Flash from './components/Flash'
import {refresh, signIn} from './actions'
import {get} from './ajax'
import {TOKEN} from './constants'

class W extends Component {
  componentDidMount() {
    const {onRefresh, checkSignIn} = this.props
    onRefresh()
    checkSignIn()
  }
  render () {
    const {children} = this.props
    return <div>
      <Header />
      <div className="container">        
        <Flash msg={this.props.location.query}/>
        {children}
        <Footer />
      </div>
    </div>
  }
}

W.propTypes = {
  onRefresh: PropTypes.func.isRequired,
  checkSignIn: PropTypes.func.isRequired,
  children: PropTypes.node.isRequired
}

const M = connect(
  (state) => {
    return {}
  },
  (dispatch) => {
    return {
      onRefresh: () => {
        get('/site/info', ).then(
          (rst)=> { dispatch(refresh(rst)) }
        )
      },
      checkSignIn: () => {
        var tkn = window.sessionStorage.getItem(TOKEN)
        if(tkn){
          dispatch(signIn(tkn))
        }
      }
    }
  }
)(W)

export default M
