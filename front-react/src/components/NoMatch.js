import React from 'react'
import i18n from 'i18next'
import { browserHistory } from 'react-router'
import {Grid, Row, Col, Thumbnail, Button} from 'react-bootstrap'

const W = () => (
  <Grid>
    <Row>
      <br/>
      <Col md={6} mdOffset={3}>
        <Thumbnail src="/fail.png" alt="242x200">
          <h3>{i18n.t('no-match.title')}</h3>
          <p>{i18n.t('no-match.body')}</p>
          <p>
            <Button onClick={() => browserHistory.push('/')} bsStyle="primary">{i18n.t('no-match.go-home')}</Button>
          </p>
        </Thumbnail>
      </Col>
    </Row>
  </Grid>
)

export default W
