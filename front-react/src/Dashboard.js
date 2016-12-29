import React, { PropTypes } from 'react'
import { connect } from 'react-redux'
import {Alert, Tabs, Tab, ButtonGroup, Button} from 'react-bootstrap'
import i18n from 'i18next'
import { browserHistory } from 'react-router'

import dashboard from './engines/dashboard'

const W = React.createClass({
  getInitialState() {
    return {
      key: 0
    };
  },
  handleSelect(key) {
    this.setState({key});
  },
  render() {
    const {user, children} = this.props
    if(!user.uid){
      return (<div className="row">
        <br/>
        <Alert bsStyle='danger'>
          <h4>{i18n.t(`flashs.alert`)}</h4>
          <p>{i18n.t('please-sign-in')}</p>
        </Alert>
      </div>)
    }
    var links = dashboard(user)
    return (<div className="row">
      <br/>
      <Tabs activeKey={this.state.key} onSelect={this.handleSelect} id="dashboard-tab">
        {
          links.map((v1, i1) => {
            return (<Tab key={i1} eventKey={i1} title={i18n.t(v1.label)}>
              <br/>
              <ButtonGroup>
                {v1.links.map((v2, i2) => (<Button onClick={() => browserHistory.push(v2.href)} key={i2}>{i18n.t(v2.label)}</Button>))}
              </ButtonGroup>
            </Tab>)
          })
        }
      </Tabs>
      <br/>
      {children}
    </div>)
  }
})

W.propTypes = {
  user: PropTypes.object.isRequired,
  children: PropTypes.node.isRequired
}

const M = connect(
  (state) => {
    return {
      user: state.currentUser
    }
  },
  (dispatch) => {
    return {}
  }
)(W)

export default M
