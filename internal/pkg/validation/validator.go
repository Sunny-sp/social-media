package validation

import (
	"reflect"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var (
	Validate   *validator.Validate
	Translator ut.Translator
)

func init() {
	Validate = validator.New()

	Validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	enLocale := en.New()
	uni := ut.New(enLocale, enLocale)
	Translator, _ = uni.GetTranslator("en")

	_ = en_translations.RegisterDefaultTranslations(Validate, Translator)
}
