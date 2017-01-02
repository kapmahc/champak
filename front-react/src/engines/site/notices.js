import React from 'react'
import i18n from 'i18next'
import {Table, FormGroup,ButtonGroup, FormControl, ControlLabel, Button, Modal} from 'react-bootstrap'

import {get, post, _delete} from '../../ajax'

export const List = React.createClass({
  getInitialState(){
    return {
      items: [],
      form: {
        body: '',
        id: null,
        show: false,
      }
    }
  },
  componentDidMount(){
    this.refresh()
  },
  refresh(){
    get('/notices').then(function(rst){
      this.setState({items:rst})
    }.bind(this))
  },
  handleNew(e){
    this.setState({
      form:{id: null, body: '', show: true}
    })
  },
  handleRemove(id, e){
    if(confirm(i18n.t('are-you-sure'))){
      _delete(`/notices/${id}`).then(function(rst){
        var items = this.state.items
        this.setState({items:items.filter(obj => obj.id !==id )})
      }.bind(this))
    }
  },
  handleEdit(n, e){
    this.setState({
      form: {body:n.body, id:n.id, show:true}
    })
  },
  onHide(e){
    this.setState({
      form: {id: null, body: '', show: false}
    })
    this.refresh()
  },
  render(){
    return <div className="col-md-12">
      <h3>{i18n.t('site.notices.title')}</h3>
      <hr/>
      <Form item={this.state.form} onHide={this.onHide}/>
      <Table striped bordered condensed hover>
        <thead>
          <tr>
            <th>ID</th>
            <th>{i18n.t('attributes.body')}</th>
            <th>{i18n.t('attributes.updatedAt')}</th>
            <th>
              {i18n.t('buttons.manage')}
              <Button onClick={this.handleNew} bsStyle="primary" bsSize="xsmall">{i18n.t('buttons.new')}</Button>
            </th>
          </tr>
        </thead>
        <tbody>
          {this.state.items.map((v,i)=>{
            return <tr key={i}>
              <td>{v.id}</td>
              <td>{v.body}</td>
              <td>{v.createdAt}</td>
              <td>
                <ButtonGroup bsSize="small">
                  <Button onClick={this.handleEdit.bind(this, v)} bsStyle="warning">{i18n.t('buttons.edit')}</Button>
                  <Button onClick={this.handleRemove.bind(this, v.id)} bsStyle="danger">{i18n.t('buttons.remove')}</Button>
                </ButtonGroup>
              </td>
            </tr>
          })}
        </tbody>
      </Table>
    </div>
  }
})

const New = React.createClass({
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
    post('/notices', data)
      .then(function(rst){
        alert(i18n.t('success'))
        // this.setState({body:''})
        browserHistory.push('/notices')        
      }.bind(this))
      .catch((err) => {
        alert(err)
      })
  },
  render(){
    const {onHide, item} = this.props
    return <Modal show={item.show} onHide={onHide} bsSize="large">
      <form onSubmit={this.handleSubmit}>
        <Modal.Header closeButton>
          <Modal.Title>{i18n.t(`buttons.${item.id ? 'edit' : 'new'}`)}</Modal.Title>
        </Modal.Header>
        <Modal.Body>
          <FormGroup controlId="body">
            <ControlLabel>{i18n.t('attributes.body')}</ControlLabel>
            <FormControl componentClass="textarea" rows={8} value={this.state.body} onChange={this.handleChange}/>
          </FormGroup>
        </Modal.Body>
        <Modal.Footer>
          <Button type="submit" bsStyle="primary">{i18n.t('buttons.submit')}</Button>
          <Button onClick={onHide}>{i18n.t('buttons.close')}</Button>
        </Modal.Footer>
      </form>
    </Modal>
  }
})
