package web

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
	"github.com/urfave/cli"
	"golang.org/x/text/language"
)

var (
	// Version version
	Version string
	// BuildTime build time
	BuildTime string
)

// Run main entry
func Run() error {

	app := cli.NewApp()
	app.Name = os.Args[0]
	app.Version = fmt.Sprintf("%s(%s)", Version, BuildTime)
	app.Usage = "CHAMPAK - A complete open source e-commerce solution by Go and React."
	app.EnableBashCompletion = true
	app.Commands = []cli.Command{}

	for _, en := range engines {
		cmd := en.Shell()
		app.Commands = append(app.Commands, cmd...)
	}

	return app.Run(os.Args)
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
		"root":  "site",
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

	viper.SetDefault("languages", []string{
		language.AmericanEnglish.String(),
		language.SimplifiedChinese.String(),
		language.TraditionalChinese.String(),
	})
}
