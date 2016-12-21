package ops

import (
	"golang.org/x/text/language"
	gin "gopkg.in/gin-gonic/gin.v1"

	"github.com/facebookgo/inject"
	"github.com/kapmahc/champak/engines/auth"
	"github.com/kapmahc/champak/web"
	"github.com/spf13/viper"
)

// Engine ops engine
type Engine struct {
	Cache    *web.Cache    `inject:""`
	Job      *web.Job      `inject:""`
	I18n     *web.I18n     `inject:""`
	Settings *web.Settings `inject:""`
	Layout   *web.Layout   `inject:""`
	Jwt      *auth.Jwt     `inject:""`
	Session  *auth.Session `inject:""`
}

// Map inject objects
func (p *Engine) Map(inj *inject.Graph) error {

	return nil
}

// Worker background workers
func (p *Engine) Worker() {}

// Dashboard dashboard links
func (p *Engine) Dashboard() web.DashboardHandler {
	return func(*gin.Context) []web.Link {
		return []web.Link{}
	}
}

// -----------------------------------------------------------------------------

func init() {
	web.Register(&Engine{})

	viper.SetEnvPrefix("champak")
	viper.BindEnv("env")
	viper.SetDefault("env", "development")

	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	viper.SetDefault("app", map[string]interface{}{
		"name": "champak",
	})

	viper.SetDefault("redis", map[string]interface{}{
		"host": "localhost",
		"port": 6379,
		"db":   8,
	})

	viper.SetDefault("database", map[string]interface{}{
		"driver": "postgres",
		"args": map[string]interface{}{
			"host":    "localhost",
			"port":    5432,
			"user":    "postgres",
			"dbname":  "champak_dev",
			"sslmode": "disable",
		},
		"pool": map[string]int{
			"max_open": 180,
			"max_idle": 6,
		},
	})

	viper.SetDefault("server", map[string]interface{}{
		"port":  8080,
		"name":  "www.change-me.com",
		"theme": "bootstrap4",
	})

	viper.SetDefault("secrets", map[string]interface{}{
		"jwt":     web.RandomStr(32),
		"aes":     web.RandomStr(32),
		"hmac":    web.RandomStr(32),
		"session": web.RandomStr(32),
		"csrf":    web.RandomStr(32),
	})

	viper.SetDefault("elasticsearch", map[string]interface{}{
		"host": "localhost",
		"port": 9200,
	})

	viper.SetDefault("rabbitmq", map[string]interface{}{
		"host":     "localhost",
		"port":     5672,
		"user":     "guest",
		"password": "guest",
		"virtual":  "",
	})

	viper.SetDefault("languages", []string{
		language.AmericanEnglish.String(),
		language.SimplifiedChinese.String(),
		language.TraditionalChinese.String(),
	})

}
