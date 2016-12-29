import React, { PropTypes } from 'react'
import { connect } from 'react-redux'
import {Navbar, NavItem, Nav} from 'react-bootstrap'
import {Link} from 'react-router'
import {IndexLinkContainer, LinkContainer} from 'react-router-bootstrap'
import i18n from 'i18next'

import LanguageBar from './LanguageBar'
import PersonalBar from './PersonalBar'

const W = ({info, user}) => (
  <Navbar inverse collapseOnSelect fixedTop fluid>
    <Navbar.Header>
      <Navbar.Brand>
        <Link to="/">{info.sub_title}</Link>
      </Navbar.Brand>
      <Navbar.Toggle />
    </Navbar.Header>
    <Navbar.Collapse>
      <Nav>
        <IndexLinkContainer to="/">
          <NavItem>{i18n.t('header.home')}</NavItem>
        </IndexLinkContainer>

        {info.top.map((v,i) => (
          <LinkContainer key={i} to={v.href}>
            <NavItem>{i18n.t(v.label)}</NavItem>
          </LinkContainer>
        ))}
        {user.uid ? <LinkContainer to="/users/logs"><NavItem>{i18n.t('header.dashboard')}</NavItem></LinkContainer> :<NavItem/>}
      </Nav>
      <Nav pullRight>
        <LanguageBar/>
        <PersonalBar/>
      </Nav>
    </Navbar.Collapse>
  </Navbar>
)

W.propTypes = {
  info: PropTypes.object.isRequired,
  user: PropTypes.object.isRequired
}

const M = connect(
  (state) => {
    return {
      info: state.siteInfo,
      user: state.currentUser
    }
  },
  (dispatch) => { return {} },
  null,
  // https://github.com/reactjs/react-redux/blob/master/docs/troubleshooting.md
  // fix nav-bar active class
  {pure: false}
)(W)

export default M
