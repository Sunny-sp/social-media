// validation/query.go
package validation

import (
	"net/http"
	"reflect"
	"strconv"
)

// BindQuery fills a struct from query parameters using `form:"..."` tags.
// Supports embedded structs (recursively).
func BindQuery(r *http.Request, dto any) error {
	values := r.URL.Query()
	return bind(values, reflect.ValueOf(dto).Elem())
}

func bind(values map[string][]string, v reflect.Value) error {
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		sf := t.Field(i)

		// Recurse into nested struct (not only anonymous)
		if field.Kind() == reflect.Struct && sf.Type.Kind() == reflect.Struct {
			if err := bind(values, field); err != nil {
				return err
			}
			continue
		}

		tag := sf.Tag.Get("form")
		if tag == "" || tag == "-" {
			continue
		}

		vals, ok := values[tag]
		if !ok || len(vals) == 0 {
			continue
		}
		val := vals[0]

		switch field.Kind() {
		case reflect.Int:
			parsed, err := strconv.Atoi(val)
			if err != nil {
				return err
			}
			field.SetInt(int64(parsed))
		case reflect.String:
			field.SetString(val)
		case reflect.Bool:
			field.SetBool(val == "true" || val == "1")
		}
	}
	return nil
}

// call this after bind()
func applyDefaults(v reflect.Value) {
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		sf := t.Field(i)

		// Recurse into nested structs
		if field.Kind() == reflect.Struct {
			applyDefaults(field)
			continue
		}

		// Skip unexported fields
		if !field.CanSet() {
			continue
		}

		// Only set default if zero value
		if !isZeroValue(field) {
			continue
		}

		def := sf.Tag.Get("default")
		if def == "" {
			continue
		}

		switch field.Kind() {
		case reflect.Int:
			if val, err := strconv.Atoi(def); err == nil {
				field.SetInt(int64(val))
			}
		case reflect.String:
			field.SetString(def)
		case reflect.Bool:
			if def == "true" || def == "1" {
				field.SetBool(true)
			} else {
				field.SetBool(false)
			}
		}
	}
}

func isZeroValue(v reflect.Value) bool {
	return reflect.DeepEqual(v.Interface(), reflect.Zero(v.Type()).Interface())
}
