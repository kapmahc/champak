package ops

import (
	"golang.org/x/text/language"
	"golang.org/x/tools/blog/atom"
	gin "gopkg.in/gin-gonic/gin.v1"

	"github.com/facebookgo/inject"
	"github.com/ikeikeikeike/go-sitemap-generator/stm"
	"github.com/jinzhu/gorm"
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
	Dao      *auth.Dao     `inject:""`
	Db       *gorm.DB      `inject:""`
}

// Map inject objects
func (p *Engine) Map(inj *inject.Graph) error {

	return nil
}

// Worker background workers
func (p *Engine) Worker() {}

// Dashboard dashboard links
func (p *Engine) Dashboard() web.DashboardHandler {
	return func(c *gin.Context) []web.Dropdown {
		var items []web.Dropdown
		user := c.MustGet(auth.CurrentUser).(*auth.User)
		if p.Dao.Is(user.ID, auth.RoleAdmin) {
			items = append(items, web.Dropdown{
				Label: "ops.dashboard.profile",
				Links: []*web.Link{
					&web.Link{
						Label: "ops.site.info.title",
						Href:  "/ops/site/info",
					},
					&web.Link{
						Label: "ops.site.author.title",
						Href:  "/ops/site/author",
					},
					&web.Link{
						Label: "ops.site.seo.title",
						Href:  "/ops/site/seo",
					},
					&web.Link{
						Label: "ops.site.status.title",
						Href:  "/ops/site/status",
					},
					&web.Link{
						Label: "ops.notices.title",
						Href:  "/ops/notices",
					},
					&web.Link{
						Label: "ops.leave_words.title",
						Href:  "/ops/leave_words",
					},
					&web.Link{
						Label: "ops.locales.title",
						Href:  "/ops/locales",
					},
				},
			})
		}
		return items
	}
}

// Atom atom entry
func (p *Engine) Atom() ([]*atom.Entry, error) {
	return []*atom.Entry{}, nil
}

// Sitemap sitemap entry
func (p *Engine) Sitemap() ([]stm.URL, error) {
	return []stm.URL{}, nil
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
