package utils

import (
	"database/sql/driver"
	"reflect"

	"github.com/go-playground/validator/v10"
)

func IsValid(data interface{}) error {
	var validate *validator.Validate
	validate = validator.New()

	validate.RegisterCustomTypeFunc(ValidateValuer)

	// build object for validation
	err := validate.Struct(data)
	if err != nil {
		return err
	}

	return nil
}

// ValidateValuer implements validator.CustomTypeFunc
func ValidateValuer(field reflect.Value) interface{} {

	if valuer, ok := field.Interface().(driver.Valuer); ok {
		val, err := valuer.Value()
		if err == nil {
			return val
		}
	}

	return nil
}
