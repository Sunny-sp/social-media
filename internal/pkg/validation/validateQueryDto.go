package validation

import (
	"net/http"
	"reflect"
)

func ValidateQueryDTO(r *http.Request, dto any) map[string]string {
	if err := BindQuery(r, dto); err != nil {
		return map[string]string{"query": "Invalid query parameter"}
	}

	applyDefaults(reflect.ValueOf(dto).Elem())

	return ValidateDTO(dto)
}
