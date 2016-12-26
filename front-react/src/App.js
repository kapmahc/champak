import React, {PropTypes} from 'react'

import Header from './components/Header'
import Footer from './components/Footer'

const W = ({children}) => (
  <div>
    <Header />
    <div className="container">
      {children}
    </div>
    <Footer />
  </div>
)

W.propTypes = {
  children: PropTypes.node.isRequired
}
export default W
