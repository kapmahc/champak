package ops

import (
	"crypto/aes"

	"github.com/SermoDigital/jose/crypto"
	"github.com/facebookgo/inject"
	"github.com/kapmahc/champak/web"
	"github.com/spf13/viper"
)

// Engine ops engine
type Engine struct {
	Cache *web.Cache `inject:""`
	Job   web.Job    `inject:""`
}

// Map inject objects
func (p *Engine) Map(inj *inject.Graph) error {
	db, err := OpenDatabase()
	if err != nil {
		return err
	}
	rep := OpenRedis()
	cip, err := aes.NewCipher([]byte(viper.GetString("secrets.aes")))
	if err != nil {
		return err
	}

	if err := inj.Provide(
		&inject.Object{Value: db},
		&inject.Object{Value: rep},
		&inject.Object{Value: cip},
		&inject.Object{Value: cip, Name: "aes.cip"},
		&inject.Object{Value: []byte(viper.GetString("secrets.hmac")), Name: "hmac.key"},
		&inject.Object{Value: []byte(viper.GetString("secrets.jwt")), Name: "jwt.key"},
		&inject.Object{Value: viper.GetString("app.name"), Name: "namespace"},
		&inject.Object{Value: crypto.SigningMethodHS512, Name: "jwt.method"},
	); err != nil {
		return err
	}
	return nil
}

// Worker background workers
func (p *Engine) Worker() {}

// -----------------------------------------------------------------------------

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
		"theme": "bootstrap4",
	})

	viper.SetDefault("secrets", map[string]interface{}{
		"jwt":     web.RandomStr(32),
		"aes":     web.RandomStr(32),
		"hmac":    web.RandomStr(32),
		"session": web.RandomStr(32),
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

	// ------------
	web.Register(&Engine{})
}
