import React from 'react'
import {Alert} from 'react-bootstrap'
import i18n from 'i18next'

const W = React.createClass({
  getInitialState() {
    return {
      show: false,
      style: '',
      subject: '',
      body: '',
    };
  },
  componentDidMount() {
    const {msg} = this.props
    if(msg.alert){
      this.setState({
        body: msg.alert,
        subject: i18n.t('flashs.alert'),
        style: 'danger',
        show: true
      })
    }else if(msg.notice){
      this.setState({
        body: msg.notice,
        subject: i18n.t('flashs.notice'),
        style: 'success',
        show: true
      })
    }
  },
  render() {    
    if(this.state.show){
      return (<div className="row">
        <br/>
        <Alert bsStyle={this.state.style} onDismiss={this.handleDismiss}>
          <h4>{this.state.subject}</h4>
          <p>{this.state.body}</p>
        </Alert>
      </div>)
    }
    return <div className="row"/>
  },
  handleDismiss() {
    this.setState({show: false});
  }
});

export default W
