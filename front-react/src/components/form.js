import React from 'react'
import {FormGroup, Button} from 'react-bootstrap'
import i18n from 'i18next'

export const Buttons = () => (
  <FormGroup>
    <Button type="submit">{i18n.t('buttons.submit')}</Button>
    &nbsp; &nbsp;
    <Button type="reset">{i18n.t('buttons.reset')}</Button>
  </FormGroup>
)
