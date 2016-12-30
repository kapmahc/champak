import React from 'react'
import i18n from 'i18next'
import {Table, Button} from 'react-bootstrap'

import {get, _delete} from '../../ajax'

export const List = React.createClass({
  getInitialState(){
    return {
      items: []
    }
  },
  componentDidMount(){
    get('/leave-words').then(function(rst){
      this.setState({items:rst})
    }.bind(this))
  },
  handleRemove(id, e){
    if(confirm(i18n.t('are-you-sure'))){
      _delete(`/leave-words/${id}`).then(function(rst){
        var items = this.state.items
        console.log(id)
        this.setState({items:items.filter(obj => obj.id !==id )})        
      }.bind(this))
    }
  },
  render(){
    return <div className="col-md-12">
      <h1>{i18n.t('site.leave-words.index.title')}</h1>
      <hr/>
      <Table striped bordered condensed hover>
        <thead>
          <tr>
            <th>ID</th>
            <th>{i18n.t('attributes.body')}</th>
            <th>{i18n.t('attributes.createdAt')}</th>
            <th>{i18n.t('buttons.manage')}</th>
          </tr>
        </thead>
        <tbody>
          {this.state.items.map((v,i)=>{
            return <tr key={i}>
              <td>{v.id}</td>
              <td>{v.body}</td>
              <td>{v.createdAt}</td>
              <td><Button onClick={this.handleRemove.bind(this, v.id)} bsStyle="danger" bsSize="small">{i18n.t('buttons.remove')}</Button></td>
            </tr>
          })}
        </tbody>
      </Table>
    </div>
  }
})
