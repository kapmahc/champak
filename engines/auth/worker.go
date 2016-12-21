package auth

import (
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/kapmahc/champak/web"
	"github.com/ugorji/go/codec"
)

// Worker background job
func (p *Engine) Worker() {
	p.Job.Receive(EmailQueue, func(mid string, body []byte, created time.Time) error {
		var mh codec.MsgpackHandle
		dec := codec.NewDecoderBytes(body, &mh)
		var args map[string]string
		if err := dec.Decode(&args); err != nil {
			return err
		}
		if web.IsProduction() {
			// TODO send mail
		} else {
			log.Debugf("send mail to %s: %s\n%s", args["to"], args["subject"], args["body"])
		}
		return nil
	})
}
