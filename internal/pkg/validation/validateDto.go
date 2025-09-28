package validation

import (
	"log"

	"github.com/go-playground/validator/v10"
)

func ValidateDTO(dto any) map[string]string {
	if err := Validate.Struct(dto); err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			messages := make(map[string]string)
			for _, e := range errs {
				// Translate and show error
				messages[e.Field()] = e.Translate(Translator)
				// Log the original error message for debugging
				log.Printf("Validation error: Field=%s, Error=%s", e.Field(), e.ActualTag())
			}
			return messages
		}
	}
	return nil
}
