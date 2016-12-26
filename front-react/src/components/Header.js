import React, { PropTypes } from 'react'
import { connect } from 'react-redux'
import {Navbar, NavItem, MenuItem, Nav, NavDropdown} from 'react-bootstrap'

import LanguageBar from './LanguageBar'
import PersonalBar from './PersonalBar'

const W = ({info}) => (
  <Navbar inverse collapseOnSelect fixedTop fluid>
    <Navbar.Header>
      <Navbar.Brand>
        <a href="#">React-Bootstrap</a>
      </Navbar.Brand>
      <Navbar.Toggle />
    </Navbar.Header>
    <Navbar.Collapse>
      <Nav>
        <NavItem eventKey={1} href="#">Link</NavItem>
        <NavItem eventKey={2} href="#">Link</NavItem>
        <NavDropdown eventKey={3} title="Dropdown" id="basic-nav-dropdown">
          <MenuItem eventKey={3.1}>Action</MenuItem>
          <MenuItem eventKey={3.2}>Another action</MenuItem>
          <MenuItem eventKey={3.3}>Something else here</MenuItem>
          <MenuItem divider />
          <MenuItem eventKey={3.3}>Separated link</MenuItem>
        </NavDropdown>
      </Nav>
      <Nav pullRight>
        <LanguageBar/>
        <PersonalBar/>
      </Nav>
    </Navbar.Collapse>
  </Navbar>
)

W.propTypes = {
  info: PropTypes.object.isRequired
}

const M = connect(
  (state) => {
    return {info:state.siteInfo}
  },
  (dispatch) => { return {} }
)(W)

export default M
