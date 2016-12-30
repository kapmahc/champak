import React, {PropTypes} from 'react'
import { connect } from 'react-redux'
import i18n from 'i18next'
import {FormGroup, HelpBlock, Button, ControlLabel, FormControl} from 'react-bootstrap'
import { Link, browserHistory } from 'react-router'

import {post} from '../../ajax'
import {signIn} from '../../actions'
import {TOKEN} from '../../constants'


export const LeaveWord = React.createClass({
  getInitialState(){
    return {
      body: ''
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
    data.append('body', this.state.body)
    post('/leave-words', data)
      .then(function(rst){
        alert(i18n.t('success'))
        this.setState({body:''})
      }.bind(this))
      .catch((err) => {
        alert(err)
      })
  },
  render () {
    return <div className="row">
      <h2>{i18n.t('site.leave-words.new.title')}</h2>
      <hr/>
      <form onSubmit={this.handleSubmit}>
        <FormGroup controlId="body">
          <ControlLabel>{i18n.t('attributes.body')}</ControlLabel>
          <FormControl componentClass="textarea" rows={8} value={this.state.body} onChange={this.handleChange}/>
        </FormGroup>
        <Button type="submit" bsStyle="primary">{i18n.t('buttons.submit')}</Button>
      </form>
      <br/>
      <SharedLinks/>
    </div>
  }
})

// ---------------------------

export const SignInW = React.createClass({
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
  handleSubmit(e) {
    e.preventDefault();
    const {onSignIn} = this.props
    var data = new FormData()
    data.append('email', this.state.email)
    data.append('password', this.state.password)
    onSignIn(data)
  },
  render (){
    return <div className="row">
      <h2>{i18n.t('auth.users.sign-in.title')}</h2>
      <hr/>
      <form onSubmit={this.handleSubmit}>
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

SignInW.propTypes = {
  onSignIn: PropTypes.func.isRequired
}

export const SignIn = connect(
  (state) => {
    return {}
  },
  (dispatch) => {
    return {
      onSignIn: (data) => {
        post('/users/sign-in', data)
          .then((rst)=>{
            window.sessionStorage.setItem(TOKEN, rst.token)
            dispatch(signIn(rst.token))
            browserHistory.push('/')
          })
          .catch((err) => {
            alert(err)
          })
      }
    }
  }
)(SignInW)

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
    <li><Link to={`/leave-words/new`}>{i18n.t('site.leave-words.new.title')}</Link></li>
  </ul>
)
