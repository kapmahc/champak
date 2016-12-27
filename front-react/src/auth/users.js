import React from 'react'
import i18n from 'i18next'
import {FormGroup, HelpBlock, Button, ControlLabel, FormControl} from 'react-bootstrap'
import { Link, browserHistory } from 'react-router'

import {post} from '../ajax'

export const SignIn = React.createClass({
  getInitialState() {
    return {
      email: '',
      password: '',
    }
  },
  handleChange(e) {
    var data = {}
    data[e.target.id] = e.target.value
    this.setState(data);
  },
  render (){
    return <div className="row">
      <h2>{i18n.t('auth.users.sign-in.title')}</h2>
      <hr/>
      <form>
        <FormGroup controlId="email">
          <ControlLabel>{i18n.t('attributes.email')}</ControlLabel>
          <FormControl type="email" value={this.state.email} onChange={this.handleChange} />
          <FormControl.Feedback />
        </FormGroup>
        <FormGroup controlId="password">
          <ControlLabel>{i18n.t('attributes.password')}</ControlLabel>
          <FormControl type="password" value={this.state.password} onChange={this.handleChange} />
          <FormControl.Feedback />
          <HelpBlock>{i18n.t('helps.password')}</HelpBlock>
        </FormGroup>
        <Button type="submit" bsStyle="primary">{i18n.t('buttons.submit')}</Button>
      </form>
      <br/>
      <SharedLinks/>
    </div>
  }
})
// -----------------------------------------------------------------------------

export const ResetPassword = React.createClass({
  getInitialState() {
    return {
      password: '',
      passwordConfirmation: '',
    }
  },
  handleChange(e) {
    var data = {}
    data[e.target.id] = e.target.value
    this.setState(data);
  },
  handleSubmit(e) {
    e.preventDefault();
    var data = new FormData()
    data.append('token', this.props.params.token)
    data.append('password', this.state.password)
    data.append('passwordConfirmation', this.state.passwordConfirmation)    
    post('/users/reset-password', data)
      .then((rst)=>{
        alert(rst.message)
        browserHistory.push('/users/sign-in')
      })
      .catch((err) => {
        alert(err)
      })
  },
  render (){
    return <div className="row">
      <h2>{i18n.t('auth.users.reset-password.title')}</h2>
      <hr/>
      <form onSubmit={this.handleSubmit}>
        <FormGroup controlId="password">
          <ControlLabel>{i18n.t('attributes.password')}</ControlLabel>
          <FormControl type="password" value={this.state.password} onChange={this.handleChange} />
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
      <br/>
      <SharedLinks/>
    </div>
  }
})

// -----------------------------------------------------------------------------
export const Confirm = ()=> (<EmailForm act='confirm' />)
export const Unlock = ()=> (<EmailForm act='unlock' />)
export const ForgotPassword = ()=> (<EmailForm act='forgot-password' />)
// -----------------------------------------------------------------------------

const EmailForm = React.createClass({
  getInitialState() {
    return {
      email: '',
    }
  },
  handleChange(e) {
    var data = {}
    data[e.target.id] = e.target.value
    this.setState(data);
  },
  handleSubmit(e) {
    e.preventDefault()
    const {act} = this.props
    var data = new FormData()
    data.append('email', this.state.email)
    post(`/users/${act}`, data)
      .then((rst)=>{
        alert(rst.message)
        browserHistory.push('/users/sign-in')
      })
      .catch((err) => {
        alert(err)
      })
  },
  render (){
    const {act} = this.props
    return <div className="row">
      <h2>{i18n.t(`auth.users.${act}.title`)}</h2>
      <hr/>
      <form onSubmit={this.handleSubmit}>
        <FormGroup controlId="email">
          <ControlLabel>{i18n.t('attributes.email')}</ControlLabel>
          <FormControl type="email" value={this.state.email} onChange={this.handleChange} />
          <FormControl.Feedback />
        </FormGroup>
        <Button type="submit" bsStyle="primary">{i18n.t('buttons.submit')}</Button>
      </form>
      <br/>
      <SharedLinks/>
    </div>
  }
})

// -----------------------------------------------------------------------------

export const SignUp = React.createClass({
  getInitialState() {
    return {
      fullName: '',
      email: '',
      password: '',
      passwordConfirmation: '',
    }
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
    data.append('email', this.state.email)
    data.append('password', this.state.password)
    data.append('passwordConfirmation', this.state.passwordConfirmation)
    post('/users/sign-up', data)
      .then((rst)=>{
        alert(rst.message)
        browserHistory.push('/users/sign-in')
      })
      .catch((err) => {
        alert(err)
      })
  },
  render (){
    return <div className="row">
      <h2>{i18n.t('auth.users.sign-up.title')}</h2>
      <hr/>
      <form onSubmit={this.handleSubmit}>
        <FormGroup controlId="fullName">
          <ControlLabel>{i18n.t('attributes.fullName')}</ControlLabel>
          <FormControl type="text" value={this.state.fullName} onChange={this.handleChange} />
          <FormControl.Feedback />
        </FormGroup>
        <FormGroup controlId="email">
          <ControlLabel>{i18n.t('attributes.email')}</ControlLabel>
          <FormControl type="email" value={this.state.email} onChange={this.handleChange} />
          <FormControl.Feedback />
        </FormGroup>
        <FormGroup controlId="password">
          <ControlLabel>{i18n.t('attributes.password')}</ControlLabel>
          <FormControl type="password" value={this.state.password} onChange={this.handleChange} />
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
      <br/>
      <SharedLinks/>
    </div>
  }
})

// -----------------------------------------------------------------------------

const SharedLinks = () => (
  <ul>
    {['sign-in', 'sign-up', 'forgot-password', 'confirm', 'unlock'].map((v, i)=>(
      <li key={i}>
        <Link to={`/users/${v}`}>{i18n.t(`auth.users.${v}.title`)}</Link>
      </li>
    ))}
  </ul>
)
