package ops

import (
	"log/syslog"

	log "github.com/Sirupsen/logrus"
	logrus_syslog "github.com/Sirupsen/logrus/hooks/syslog"
	"github.com/kapmahc/champak/web"
	"github.com/spf13/viper"
)

func init() {
	if web.IsProduction() {
		log.SetLevel(log.InfoLevel)
		if wrt, err := syslog.New(syslog.LOG_INFO, viper.GetString("app.name")); err == nil {
			log.AddHook(&logrus_syslog.SyslogHook{Writer: wrt})
		} else {
			log.Error(err)
		}
	} else {
		log.SetLevel(log.DebugLevel)
	}
}
