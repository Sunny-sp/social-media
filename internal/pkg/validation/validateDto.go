package validation

import "github.com/go-playground/validator/v10"

func ValidateDTO(dto any) map[string]string {
	if err := Validate.Struct(dto); err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			messages := make(map[string]string)
			for _, e := range errs {
				messages[e.Field()] = e.Translate(Translator)
			}
			return messages
		}
	}
	return nil
}
