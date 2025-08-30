package validation

import (
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

	enLocale := en.New()
	uni := ut.New(enLocale, enLocale)
	Translator, _ = uni.GetTranslator("en")

	_ = en_translations.RegisterDefaultTranslations(Validate, Translator)
}
