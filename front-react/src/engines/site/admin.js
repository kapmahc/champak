import React, {PropTypes} from 'react'
import { connect } from 'react-redux'
import {ListGroupItem, ListGroup, Checkbox, FormGroup, FormControl, ControlLabel, Button, HelpBlock} from 'react-bootstrap'
import i18n from 'i18next'

import {get, post} from '../../ajax'

export const Status = React.createClass({
  getInitialState(){
    return {
      cache: {},
      jobs:{},
      os:{},
      redis: '',
      db:{},
    }
  },
  componentDidMount(){
    get('/admin/site/status').then(function(rst){
      this.setState(rst)
    }.bind(this))
  },
  render(){
    const {os, db, redis} = this.state
    return  <div className="row">
      <div className="col-md-4">
        <h3>{i18n.t('site.status.os')}</h3>
        <hr/>
        <ListGroup>
          {Object.keys(os).map((k,i)=>(<ListGroupItem key={i}>{k}: {os[k]}</ListGroupItem>))}
        </ListGroup>
      </div>
      <div className="col-md-4">
        <h3>{i18n.t('site.status.database')}</h3>
        <hr/>
        <ListGroup>
          {Object.keys(db).map((k,i)=>(<ListGroupItem key={i}>{k}: {db[k]}</ListGroupItem>))}
        </ListGroup>
      </div>
      <div className="col-md-4">
        <h3>REDIS</h3>
        <hr/>
        <pre><code>{redis}</code></pre>
      </div>
      <div className="col-md-4">
        <h3>{i18n.t('site.status.cache')}</h3>
        <hr/>
      </div>
      <div className="col-md-4">
        <h3>{i18n.t('site.status.jobs')}</h3>
        <hr/>
      </div>
    </div>
  }
})

// ----------------------

export const Smtp = React.createClass({
  getInitialState() {
    return {
      host: '',
      port: 25,
      user: '',
      password: '',
      passwordConfirmation: '',
      ssl: false,
    };
  },
  componentDidMount(){
    get('/admin/site/smtp').then(
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
  handleSslChecked(e){
    this.setState({ssl: e.target.checked})
  },
  handleSubmit(e) {
    e.preventDefault();
    var data = new FormData()
    data.append('host', this.state.host)
    data.append('port', this.state.port)
    data.append('user', this.state.user)
    data.append('password', this.state.password)
    data.append('passwordConfirmation', this.state.passwordConfirmation)
    data.append('ssl', this.state.ssl)
    post('/admin/site/smtp', data)
      .then((rst)=>{
        alert(i18n.t('success'))
      })
      .catch((err) => {
        alert(err)
      })
  },
  render(){
    return <div className="col-md-12">
      <h3>{i18n.t('site.smtp.title')}</h3>
      <hr/>
      <form onSubmit={this.handleSubmit}>
        <FormGroup controlId="host">
          <ControlLabel>{i18n.t('attributes.host')}</ControlLabel>
          <FormControl type="text" value={this.state.host} onChange={this.handleChange} />
          <FormControl.Feedback />
        </FormGroup>
        <FormGroup controlId="port">
          <ControlLabel>{i18n.t('attributes.port')}</ControlLabel>
          <FormControl componentClass="select" value={this.state.port} onChange={this.handleChange}>
            {[25, 465, 587].map((v)=>(<option key={v} value={v}>{v}</option>))}
          </FormControl>
        </FormGroup>
        <Checkbox checked={this.state.ssl} onChange={this.handleSslChecked}>
          {i18n.t('attributes.ssl')}
        </Checkbox>
        <FormGroup controlId="user">
          <ControlLabel>{i18n.t('attributes.user')}</ControlLabel>
          <FormControl type="email" value={this.state.user} onChange={this.handleChange} />
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
    </div>
  }
})

// --------------------------------------

const SeoW = React.createClass({
  getInitialState() {
    return {
      google: '',
      baidu: '',
    };
  },
  componentDidMount(){
    get('/admin/site/seo').then(
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
    data.append('google', this.state.google)
    data.append('baidu', this.state.baidu)
    post('/admin/site/seo', data)
      .then((rst)=>{
        alert(i18n.t('success'))
      })
      .catch((err) => {
        alert(err)
      })
  },
  render(){
    const {info} = this.props
    return <div className="col-md-12">
      <h3>{i18n.t('site.seo.title')}</h3>
      <hr/>
      <form onSubmit={this.handleSubmit}>
        <FormGroup controlId="google">
          <ControlLabel>{i18n.t('site.attributes.googleVerifyID')}</ControlLabel>
          <FormControl type="text" value={this.state.google} onChange={this.handleChange} />
          <FormControl.Feedback />
        </FormGroup>
        <FormGroup controlId="baidu">
          <ControlLabel>{i18n.t('site.attributes.baiduVerifyID')}</ControlLabel>
          <FormControl type="text" value={this.state.baidu} onChange={this.handleChange} />
          <FormControl.Feedback />
        </FormGroup>
        <Button type="submit" bsStyle="primary">{i18n.t('buttons.submit')}</Button>
      </form>
      <br/>
      <ListGroup>
        <ListGroupItem><a target="_blank" href="/robots.txt">robots.txt</a></ListGroupItem>
        <ListGroupItem><a target="_blank" href="/sitemap.xml.gz">sitemap.xml.gz</a></ListGroupItem>
        {info.languages.map((v,i)=>(<ListGroupItem key={i}><a target="_blank" href="/rss/{v}.atom">{v}.atom</a></ListGroupItem>))}
      </ListGroup>
    </div>
  }
})


SeoW.propTypes = {
  info: PropTypes.object.isRequired
}

export const Seo = connect(
  (state) => {
    return {info: state.siteInfo}
  },
  (dispatch) => {
    return {}
  }
)(SeoW)

// --------------------------

export const Info = React.createClass({
  getInitialState() {
    return {
      title: '',
      subTitle: '',
      keywords: '',
      description: '',
      copyright: '',
    };
  },
  componentDidMount(){
    get('/admin/site/info').then(
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
    data.append('title', this.state.title)
    data.append('subTitle', this.state.subTitle)
    data.append('keywords', this.state.keywords)
    data.append('description', this.state.description)
    data.append('copyright', this.state.copyright)
    post('/admin/site/info', data)
      .then((rst)=>{
        alert(i18n.t('success'))
      })
      .catch((err) => {
        alert(err)
      })
  },
  render(){
    return <div className="col-md-12">
      <h3>{i18n.t('site.info.title')}</h3>
      <hr/>
      <form onSubmit={this.handleSubmit}>
        <FormGroup controlId="title">
          <ControlLabel>{i18n.t('site.attributes.title')}</ControlLabel>
          <FormControl type="text" value={this.state.title} onChange={this.handleChange} />
          <FormControl.Feedback />
        </FormGroup>
        <FormGroup controlId="subTitle">
          <ControlLabel>{i18n.t('site.attributes.subTitle')}</ControlLabel>
          <FormControl type="text" value={this.state.subTitle} onChange={this.handleChange} />
          <FormControl.Feedback />
        </FormGroup>
        <FormGroup controlId="keywords">
          <ControlLabel>{i18n.t('site.attributes.keywords')}</ControlLabel>
          <FormControl type="text" value={this.state.keywords} onChange={this.handleChange} />
          <FormControl.Feedback />
        </FormGroup>
        <FormGroup controlId="description">
         <ControlLabel>{i18n.t('site.attributes.description')}</ControlLabel>
         <FormControl componentClass="textarea" value={this.state.description} onChange={this.handleChange} />
       </FormGroup>
        <FormGroup controlId="copyright">
          <ControlLabel>{i18n.t('site.attributes.copyright')}</ControlLabel>
          <FormControl type="text" value={this.state.copyright} onChange={this.handleChange} />
          <FormControl.Feedback />
        </FormGroup>
        <Button type="submit" bsStyle="primary">{i18n.t('buttons.submit')}</Button>
      </form>
    </div>
  }
})

// ------------------------------

export const Author = React.createClass({
  getInitialState() {
    return {
      email: '',
      name: '',
    };
  },
  componentDidMount(){
    get('/admin/site/author').then(
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
    data.append('name', this.state.name)
    data.append('email', this.state.email)
    post('/admin/site/author', data)
      .then((rst)=>{
        alert(i18n.t('success'))
      })
      .catch((err) => {
        alert(err)
      })
  },
  render(){
    return <div className="col-md-12">
      <h3>{i18n.t('site.author.title')}</h3>
      <hr/>
      <form onSubmit={this.handleSubmit}>
        <FormGroup controlId="name">
          <ControlLabel>{i18n.t('site.attributes.author.name')}</ControlLabel>
          <FormControl type="text" value={this.state.name} onChange={this.handleChange} />
          <FormControl.Feedback />
        </FormGroup>
        <FormGroup controlId="email">
          <ControlLabel>{i18n.t('site.attributes.author.email')}</ControlLabel>
          <FormControl type="email" value={this.state.email} onChange={this.handleChange} />
          <FormControl.Feedback />
        </FormGroup>
        <Button type="submit" bsStyle="primary">{i18n.t('buttons.submit')}</Button>
      </form>
    </div>
  }
})
