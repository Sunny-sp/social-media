package validation

import (
	"log"
	"path/filepath"
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

func allowedFileExtension(fl validator.FieldLevel) bool {
	filename := fl.Field().String()

	if filename == "" {
		log.Println("Filename is empty, skipping extension check")
		return true // Let 'required' handle empty cases
	}

	// Get the file extension with dot (e.g., ".jpeg")
	ext := strings.ToLower(filepath.Ext(filename))

	// Get the allowed extensions from the tag parameter
	param := fl.Param()
	if param == "" {
		log.Println("No allowed extensions specified in tag")
		return false
	}

	// Split on spaces and clean up
	allowed := strings.Fields(param) // Use Fields to handle multiple spaces

	for _, a := range allowed {
		a = strings.ToLower(strings.TrimSpace(a))
		// Normalize to include dot for comparison
		if !strings.HasPrefix(a, ".") {
			a = "." + a
		}

		if ext == a {
			log.Printf("Extension %s is allowed", ext)
			return true
		}
	}

	return false
}

func init() {
	Validate = validator.New()

	// Register custom tag name function to map JSON field names
	Validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	// Register the custom validation function
	err := Validate.RegisterValidation("allowed_ext", allowedFileExtension)
	if err != nil {
		log.Fatalf("Failed to register allowed_ext validation: %v", err)
	}

	// Setup translator for English
	enLocale := en.New()
	uni := ut.New(enLocale, enLocale)
	Translator, _ = uni.GetTranslator("en")

	// Register default translations
	err = en_translations.RegisterDefaultTranslations(Validate, Translator)
	if err != nil {
		log.Fatalf("Failed to register default translations: %v", err)
	}

	// Register custom translation for allowed_ext
	err = Validate.RegisterTranslation("allowed_ext", Translator, func(ut ut.Translator) error {
		return ut.Add("allowed_ext", "Only these file types are allowed: {0}", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		// Split on spaces and format as "jpeg, jpg, png or mp4"
		params := strings.Fields(fe.Param())
		if len(params) == 0 {
			return "No valid file types specified"
		}
		for i, p := range params {
			params[i] = strings.TrimSpace(p)
		}
		formattedParams := strings.Join(params[:len(params)-1], ", ") + " or " + params[len(params)-1]
		t, _ := ut.T("allowed_ext", formattedParams)
		return t
	})
	if err != nil {
		log.Fatalf("Failed to register translation for allowed_ext: %v", err)
	}
}
