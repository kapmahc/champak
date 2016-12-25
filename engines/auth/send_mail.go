package auth

func (p *Engine) sendEmail(lng string, user *User, act string) {

}

func (p *Engine) parseToken(token string, act string) (*User, error) {
	cm, err := p.Jwt.Validate([]byte(token))
	if err != nil {
		return nil, err
	}
	return p.Dao.GetUserByUID(cm.Get("uid").(string))
}
