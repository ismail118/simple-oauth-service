package helper

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type CustomValidator struct {
}

func NewValidator() *CustomValidator {
	return &CustomValidator{}
}

func (v *CustomValidator) ValidateStruct(s interface{}) error {
	if reflect.ValueOf(s).Kind() == reflect.Struct {
		rt := reflect.TypeOf(s)
		rv := reflect.ValueOf(s)

		for i := 0; i < rt.NumField(); i++ {
			field := rt.Field(i)
			if validate, ok := field.Tag.Lookup("validate"); ok {
				validations := strings.Split(validate, ",")
				for _, validation := range validations {
					if validation == "required" {
						val := reflect.ValueOf(field)
						switch val.Kind() {
						case reflect.String:
							if valueString := rv.Field(i).Interface().(string); valueString == "" {
								return errors.New(fmt.Sprintf("%v can't be empty", field.Name))
							}
						case reflect.Int:
							if valueInt := rv.Field(i).Interface().(int); valueInt < 1 {
								return errors.New(fmt.Sprintf("%v can't be empty", field.Name))
							}
						case reflect.Int64:
							if valueInt := rv.Field(i).Interface().(int64); valueInt < 1 {
								return errors.New(fmt.Sprintf("%v can't be empty", field.Name))
							}
						}
					}
				}
			}
		}
	} else {
		return errors.New(fmt.Sprintf("%v is not struct", s))
	}

	return nil
}
