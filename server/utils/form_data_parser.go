package utils

import (
	"reflect"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func FormDataParser[T any](c *fiber.Ctx) (T, error) {
	var out T
	var prod T

	prodValue := reflect.ValueOf(&prod).Elem()
	prodTag := reflect.TypeOf(out)

	for i := 0; i < prodValue.NumField(); i++ {
		field := prodValue.Field(i)
		tag := prodTag.Field(i).Tag
		if !field.CanSet() {
			continue
		}
		var jsonTag string
		if jsonTag = tag.Get("json"); jsonTag == "" {
			continue
		}
		switch field.Kind() {
		case reflect.String:
			field.SetString(c.FormValue(jsonTag))
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			iv, err := strconv.ParseInt(c.FormValue(jsonTag), 10, field.Type().Bits())
			if err != nil {
				return out, err
			}
			field.SetInt(iv)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			uv, err := strconv.ParseUint(c.FormValue(jsonTag), 10, field.Type().Bits())
			if err != nil {
				return out, err
			}
			field.SetUint(uv)
		case reflect.Bool:
			bv, err := strconv.ParseBool(c.FormValue(jsonTag))
			if err != nil {
				return out, err
			}
			field.SetBool(bv)
		case reflect.Float32, reflect.Float64:
			field64, err := strconv.ParseFloat(c.FormValue(jsonTag), field.Type().Bits())
			if err != nil {
				return out, err
			}
			field.SetFloat(field64)
		default:
			continue
		}
	}

	return prod, nil
}
