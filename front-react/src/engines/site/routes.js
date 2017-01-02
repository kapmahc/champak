import React from 'react'
import {Route, IndexRoute} from 'react-router'

import Home from './Home'
import NoMatch from './NoMatch'
import Dashboard from '../../Dashboard'
import {Info as SiteInfo, Author as SiteAuthor, Seo as SiteSeo, Smtp as SiteSmtp, Status as SiteStatus} from './admin'
import {List as IndexLeaveWords} from './leave-words'
import {List as IndexNotices} from './notices'

const W = [
  <IndexRoute key="index" component={Home}/>,
  <Route key="home" path="home" component={Home}/>,
  <Route key="site.admin" path="admin" component={Dashboard}>
    <Route path="site/info" component={SiteInfo}/>
    <Route path="site/author" component={SiteAuthor}/>
    <Route path="site/seo" component={SiteSeo}/>
    <Route path="site/smtp" component={SiteSmtp}/>
    <Route path="site/status" component={SiteStatus}/>
  </Route>,
  <Route key="leave-words" path="leave-words" component={Dashboard}>
    <IndexRoute component={IndexLeaveWords}/>
  </Route>,
  <Route key="notices" path="notices" component={Dashboard}>
    <IndexRoute component={IndexNotices}/>
  </Route>,
  <Route key="no-match" path="*" component={NoMatch}/>
]

export default W
