package auth

import (
	"crypto/aes"

	"github.com/SermoDigital/jose/crypto"
	"github.com/facebookgo/inject"
	"github.com/kapmahc/champak/web"
	"github.com/kapmahc/champak/web/cache"
	"github.com/kapmahc/champak/web/i18n"
	"github.com/spf13/viper"
	gin "gopkg.in/gin-gonic/gin.v1"
)

// Engine  auth engine
type Engine struct {
	Cache cache.Store `inject:""`
	Job   web.Job     `inject:""`
}

// Map map objects
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

	i1n := i18n.I18n{Locales: make(map[string]map[string]string)}
	if err := inj.Provide(
		&inject.Object{Value: db},
		&inject.Object{Value: rep},
		&inject.Object{Value: &i18n.GormStore{}},
		&inject.Object{Value: &i1n},
		&inject.Object{Value: cip},
		&inject.Object{Value: cip, Name: "aes.cip"},
		&inject.Object{Value: []byte(viper.GetString("secrets.hmac")), Name: "hmac.key"},
		&inject.Object{Value: []byte(viper.GetString("secrets.jwt")), Name: "jwt.key"},
		&inject.Object{Value: viper.GetString("app.name"), Name: "namespace"},
		&inject.Object{Value: crypto.SigningMethodHS512, Name: "jwt.method"},
		&inject.Object{Value: &cache.RedisStore{}},
	); err != nil {
		return err
	}
	return nil
}

// Mount mount web points
func (p *Engine) Mount(*gin.Engine) {}

// Worker background job
func (p *Engine) Worker() {}

// -----------------------------------------------------------------------------

func init() {
	web.Register(&Engine{})
}
