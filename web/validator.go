package web

import (
	"errors"
	"net/http"
	"reflect"
	"strings"
	"sync"

	"github.com/gin-contrib/sessions"

	gin "gopkg.in/gin-gonic/gin.v1"
	"gopkg.in/gin-gonic/gin.v1/binding"
	validator "gopkg.in/go-playground/validator.v9"
)

// FlashHandler show flash
func FlashHandler(ego string, fn func(*gin.Context) error) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := fn(c); err != nil {
			ss := sessions.Default(c)
			ss.AddFlash(err.Error(), ALERT)
			ss.Save()
			c.Redirect(http.StatusInternalServerError, ego)
		}
	}
}

// PostFormHandler fix gin bind error return 400
func PostFormHandler(ego string, fm interface{}, fn func(*gin.Context, interface{}) error) gin.HandlerFunc {
	return FlashHandler(ego, func(c *gin.Context) error {
		err := binding.FormPost.Bind(c.Request, fm)
		if err == nil {
			err = fn(c, fm)
		} else {
			err = errors.New(strings.Replace(err.Error(), "\n", "<br/>", -1))
		}
		return err
	})
}

type defaultValidator struct {
	once     sync.Once
	validate *validator.Validate
}

var _ binding.StructValidator = &defaultValidator{}

func (v *defaultValidator) ValidateStruct(obj interface{}) error {
	if kindOfData(obj) == reflect.Struct {
		v.lazyinit()
		if err := v.validate.Struct(obj); err != nil {
			return err
		}
	}
	return nil
}

func (v *defaultValidator) lazyinit() {
	v.once.Do(func() {
		v.validate = validator.New()
		v.validate.SetTagName("binding")
		// add any custom validations etc. here
	})
}

func kindOfData(data interface{}) reflect.Kind {
	value := reflect.ValueOf(data)
	valueType := value.Kind()

	if valueType == reflect.Ptr {
		valueType = value.Elem().Kind()
	}
	return valueType
}

func init() {
	binding.Validator = new(defaultValidator)
}
