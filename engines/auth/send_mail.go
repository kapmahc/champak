package auth

import (
	"bytes"
	"fmt"

	"github.com/SermoDigital/jose/jws"
	"github.com/ugorji/go/codec"
)

const (
	// EmailQueue email queue name
	EmailQueue = "emails"
)

func (p *Engine) sendEmail(lng string, user *User, act string) error {
	cm := jws.Claims{}
	cm.Set("act", act)
	cm.Set("uid", user.UID)
	tkn, err := p.Jwt.Sum(cm, 1)
	if err != nil {
		return err
	}

	obj := struct {
		Href string
	}{
		Href: fmt.Sprintf("%s/personal/%s/%s", Home(), act, tkn),
	}
	subject, err := p.I18n.F(lng, fmt.Sprintf("auth.emails.%s.subject", act), obj)
	if err != nil {
		return err
	}
	body, err := p.I18n.F(lng, fmt.Sprintf("auth.emails.%s.body", act), obj)
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	var mh codec.MsgpackHandle
	enc := codec.NewEncoder(&buf, &mh)

	if err = enc.Encode(map[string]string{
		"to":      user.Email,
		"subject": subject,
		"body":    body,
	}); err != nil {
		return err
	}
	p.Job.Send(EmailQueue, buf.Bytes())
	return nil
}

func (p *Engine) parseToken(token string, act string) (*User, error) {
	cm, err := p.Jwt.Validate([]byte(token))
	if err != nil {
		return nil, err
	}
	return p.Dao.GetUserByUID(cm.Get("uid").(string))
}
