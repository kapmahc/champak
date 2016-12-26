import React from 'react'
import i18n from 'i18next'
import { Link } from 'react-router'
import {FormGroup, HelpBlock, Button, ControlLabel, FormControl} from 'react-bootstrap'

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
    post('/users/sign-up', this.state)
      .then((rst)=>{
        console.log(rst)
      })
      .catch((err) => {
        console.log(err);
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
