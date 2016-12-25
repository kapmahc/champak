package auth

import (
	"net/http"

	"github.com/facebookgo/inject"
	"github.com/ikeikeikeike/go-sitemap-generator/stm"
	"github.com/jinzhu/gorm"
	"github.com/kapmahc/champak/web"
	"github.com/kapmahc/champak/web/crypto"
	"github.com/spf13/viper"
	"github.com/unrolled/render"
	"golang.org/x/text/language"
	"golang.org/x/tools/blog/atom"
)

// Engine auth engine
type Engine struct {
	Dao                   *Dao                   `inject:""`
	Db                    *gorm.DB               `inject:""`
	Helper                *Helper                `inject:""`
	Render                *render.Render         `inject:""`
	Jwt                   *Jwt                   `inject:""`
	CurrentUserMiddleware *CurrentUserMiddleware `inject:""`
}

// Map inject objects
func (p *Engine) Map(*inject.Graph) error {
	return nil
}

// Worker register background jobs
func (p *Engine) Worker() {}

// Dashboard dashboard links
func (p *Engine) Dashboard() web.DashboardHandler {
	return func(req *http.Request) []web.Dropdown {
		return []web.Dropdown{}
	}
}

// Atom rss-atom
func (p *Engine) Atom() ([]*atom.Entry, error) {
	return []*atom.Entry{}, nil
}

// Sitemap sitemap.xml.gz
func (p *Engine) Sitemap() ([]stm.URL, error) {
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
		"jwt":     crypto.Rand(32),
		"aes":     crypto.Rand(32),
		"hmac":    crypto.Rand(32),
		"session": crypto.Rand(32),
		"csrf":    crypto.Rand(32),
	})

	viper.SetDefault("elasticsearch", map[string]interface{}{
		"host": "localhost",
		"port": 9200,
	})

	viper.SetDefault("languages", []string{
		language.AmericanEnglish.String(),
		language.SimplifiedChinese.String(),
		language.TraditionalChinese.String(),
	})

	web.Register(&Engine{})
}
