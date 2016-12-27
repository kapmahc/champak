package auth

import (
	"fmt"

	"github.com/RichardKnop/machinery/v1/signatures"
	"github.com/SermoDigital/jose/jws"
	log "github.com/Sirupsen/logrus"
	"github.com/kapmahc/champak/web"
	"github.com/spf13/viper"
)

const (
	sendEmailJob = "auth.send-email"
)

// Workers background jobs
func (p *Engine) Workers() {
	p.Server.RegisterTask(sendEmailJob, p.doSendEmail)
}

func (p *Engine) sendEmail(lng string, user *User, act string) {
	cm := jws.Claims{}
	cm.Set("act", act)
	cm.Set("uid", user.UID)
	tkn, err := p.Jwt.Sum(cm, 1)
	if err != nil {
		log.Error(err)
		return
	}

	obj := struct {
		Frontend string
		Backend  string
		Token    string
	}{
		Frontend: viper.GetString("server.frontend"),
		Backend:  viper.GetString("server.backend"),
		Token:    string(tkn),
	}
	subject, err := p.I18n.F(lng, fmt.Sprintf("auth.emails.%s.subject", act), obj)
	if err != nil {
		log.Error(err)
		return
	}
	body, err := p.I18n.F(lng, fmt.Sprintf("auth.emails.%s.body", act), obj)
	if err != nil {
		log.Error(err)
		return
	}

	// -----------------------
	task := signatures.TaskSignature{
		Name: sendEmailJob,
		Args: []signatures.TaskArg{
			signatures.TaskArg{
				Type:  "string",
				Value: user.Email,
			},
			signatures.TaskArg{
				Type:  "string",
				Value: subject,
			},
			signatures.TaskArg{
				Type:  "string",
				Value: body,
			},
		},
	}

	if _, err := p.Server.SendTask(&task); err != nil {
		log.Error(err)
	}
}

func (p *Engine) doSendEmail(to, subject, body string) (interface{}, error) {
	if !web.IsProduction() {
		log.Infof("to %s\n%s\n%s", to, subject, body)
		return "done", nil
	}
	// TODO
	return "done", nil
}
