package auth

import (
	"net/http"

	"github.com/facebookgo/inject"
	"github.com/ikeikeikeike/go-sitemap-generator/stm"
	"github.com/jinzhu/gorm"
	"github.com/kapmahc/champak/web"
	"github.com/spf13/viper"
	"github.com/unrolled/render"
	"golang.org/x/text/language"
	"golang.org/x/tools/blog/atom"
)

// Engine auth engine
type Engine struct {
	Cache    *web.Cache     `inject:""`
	Job      *web.Job       `inject:""`
	I18n     *web.I18n      `inject:""`
	Settings *web.Settings  `inject:""`
	Jwt      *Jwt           `inject:""`
	Dao      *Dao           `inject:""`
	Db       *gorm.DB       `inject:""`
	Render   *render.Render `inject:""`
}

// Map map objects
func (*Engine) Map(*inject.Graph) error {
	return nil
}

// Dashboard dashboard links(by user)
func (*Engine) Dashboard(req *http.Request) []web.Dropdown {
	return []web.Dropdown{}
}

// Atom rss-atom
func (*Engine) Atom() ([]*atom.Entry, error) {
	return []*atom.Entry{}, nil
}

// Sitemap sitemap.xml
func (*Engine) Sitemap() ([]stm.URL, error) {
	return []stm.URL{}, nil
}

func init() {
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
		"theme": "bootstrap",
	})

	viper.SetDefault("secrets", map[string]interface{}{
		"jwt":     RandomStr(32),
		"aes":     RandomStr(32),
		"hmac":    RandomStr(32),
		"session": RandomStr(32),
		"csrf":    RandomStr(32),
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

	web.Register(&Engine{})
}
