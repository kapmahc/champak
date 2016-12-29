import React from 'react'
import {ListGroupItem, ListGroup, FormGroup, FormControl, ControlLabel, Button, HelpBlock} from 'react-bootstrap'
import i18n from 'i18next'

import {get, post} from '../../ajax'

export const Info = React.createClass({
  getInitialState() {
    return {
      email: '',
      fullName: '',
      home: '',
      logo: '',
    };
  },
  componentDidMount(){
    get('/users/info').then(
      function (rst){
        this.setState(rst)
      }.bind(this)
    )
  },
  handleChange(e) {
    var data = {}
    data[e.target.id] = e.target.value
    this.setState(data);
  },
  handleSubmit(e) {
    e.preventDefault();
    var data = new FormData()
    data.append('fullName', this.state.fullName)
    data.append('home', this.state.home)
    data.append('logo', this.state.logo)
    post('/users/info', data)
      .then((rst)=>{
        alert(i18n.t('success'))
      })
      .catch((err) => {
        alert(err)
      })
  },
  render(){
    return <div className="col-md-12">
      <h3>{i18n.t('auth.users.info.title')}</h3>
      <hr/>
      <form onSubmit={this.handleSubmit}>
        <FormGroup controlId="email">
          <ControlLabel>{i18n.t('attributes.email')}</ControlLabel>
          <FormControl disabled type="email" value={this.state.email} onChange={this.handleChange} />
          <FormControl.Feedback />
        </FormGroup>
        <FormGroup controlId="fullName">
          <ControlLabel>{i18n.t('attributes.fullName')}</ControlLabel>
          <FormControl type="text" value={this.state.fullName} onChange={this.handleChange} />
          <FormControl.Feedback />
        </FormGroup>
        <FormGroup controlId="home">
          <ControlLabel>{i18n.t('auth.attributes.user.home')}</ControlLabel>
          <FormControl type="text" value={this.state.home} onChange={this.handleChange} />
          <FormControl.Feedback />
        </FormGroup>
        <FormGroup controlId="logo">
          <ControlLabel>{i18n.t('auth.attributes.user.logo')}</ControlLabel>
          <FormControl type="text" value={this.state.logo} onChange={this.handleChange} />
          <FormControl.Feedback />
        </FormGroup>
        <Button type="submit" bsStyle="primary">{i18n.t('buttons.submit')}</Button>
      </form>
    </div>
  }
})

export const ChangePassword = React.createClass({
  getInitialState() {
    return {
      password: '',
      newPassword: '',
      passwordConfirmation: '',
    };
  },
  handleChange(e) {
    var data = {}
    data[e.target.id] = e.target.value
    this.setState(data);
  },
  handleSubmit(e) {
    e.preventDefault();
    var data = new FormData()
    data.append('password', this.state.password)
    data.append('newPassword', this.state.newPassword)
    data.append('passwordConfirmation', this.state.passwordConfirmation)
    post('/users/change-password', data)
      .then((rst)=>{
        alert(rst.message)
      })
      .catch((err) => {
        alert(err)
      })
  },
  render(){
    return (<div className="col-md-12">
      <h3>{i18n.t('auth.users.change-password.title')}</h3>
      <hr/>
      <form onSubmit={this.handleSubmit}>
        <FormGroup controlId="password">
          <ControlLabel>{i18n.t('attributes.password')}</ControlLabel>
          <FormControl type="password" value={this.state.password} onChange={this.handleChange} />
          <FormControl.Feedback />
          <HelpBlock>{i18n.t('auth.helps.need-password')}</HelpBlock>
        </FormGroup>
        <FormGroup controlId="newPassword">
          <ControlLabel>{i18n.t('attributes.newPassword')}</ControlLabel>
          <FormControl type="password" value={this.state.newPassword} onChange={this.handleChange} />
          <FormControl.Feedback />
          <HelpBlock>{i18n.t('helps.password')}</HelpBlock>
        </FormGroup>
        <FormGroup controlId="passwordConfirmation">
          <ControlLabel>{i18n.t('attributes.passwordConfirmation')}</ControlLabel>
          <FormControl type="password" value={this.state.passwordConfirmation} onChange={this.handleChange} />
          <FormControl.Feedback />
          <HelpBlock>{i18n.t('helps.passwordConfirmation')}</HelpBlock>
        </FormGroup>
        <Button type="submit" bsStyle="primary">{i18n.t('buttons.submit')}</Button>
      </form>
    </div>)
  }
})

export const Logs = React.createClass({
  getInitialState() {
    return {
      items: []
    };
  },
  componentDidMount(){
    get('/users/logs').then(
      function (rst){
        this.setState({items:rst})
      }.bind(this)
    )
  },
  render(){
    return <div className="col-md-12">
      <h3>{i18n.t('auth.users.logs.title')}</h3>
      <hr/>
      <ListGroup>
        {this.state.items.map((v, i) => (<ListGroupItem key={i}>{v.createdAt}: [{v.ip}] {v.message}</ListGroupItem>))}
      </ListGroup>
    </div>
  }
})
