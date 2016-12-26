import React, {PropTypes, Component} from 'react'
import { connect } from 'react-redux'

import Header from './components/Header'
import Footer from './components/Footer'
import {refresh} from './actions'
import {get} from './ajax'

class W extends Component {
  componentDidMount() {
    const {onRefresh} = this.props
    onRefresh()
  }
  render () {
    const {children} = this.props
    return <div>
      <Header />
      <div className="container">
        {children}
      </div>
      <Footer />
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
      onRefresh: () => {
        get('/site/info').then(function(rst) {
          dispatch(refresh(rst))
        })
      }
    }
  }
)(W)

export default M
