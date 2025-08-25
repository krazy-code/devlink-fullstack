// package utils

// import (
// 	"reflect"
// 	"strings"

// 	"github.com/gofiber/fiber/v2"
// )

// // FormDataParser returns a set of field names (and common tags) for the generic type T.
// // It expects T to be a struct (or pointer to a struct).
// func FormDataParser[T any](c *fiber.Ctx) any {
// 	keyMap := make(map[string]string)

// 	var zero T
// 	t := reflect.TypeOf(zero)
// 	if t == nil {
// 		return keyMap
// 	}
// 	if t.Kind() == reflect.Ptr {
// 		t = t.Elem()
// 	}

// 	for i := 0; i < t.NumField(); i++ {
// 		f := t.Field(i)
// 		// Field name
// 		// keyMap[f.Name] = c.FormValue(name)
// 		// // form tag
// 		// if formTag := f.Tag.Get("form"); formTag != "" {
// 		// 	keyMap[formTag] = ""
// 		// }
// 		// // json tag (take name before comma)
// 		if jsonTag := f.Tag.Get("json"); jsonTag != "" {
// 			name := strings.Split(jsonTag, ",")[0]
// 			if name != "" && name != "-" {
// 				keyMap[f.Name] = c.FormValue(name)
// 			}
// 		}
// 	}

// 	return keyMap
// }

// // return T{
// // 	// Name:  c.FormValue("name"),
// // 	// Bio:   c.FormValue("bio"),
// // 	// Email: c.FormValue("email"),
// // 	// // Add more fields as needed
// // }

package utils

import (
	"reflect"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// FormDataParser returns a set of field names (and common tags) for the generic type T.
// It expects T to be a struct (or pointer to a struct).
// It still returns `any` (the concrete T value) so existing call sites don't need a signature change.
func FormDataParser[T any](c *fiber.Ctx) any {
	keyMap := make(map[string]string)

	var zero T
	t := reflect.TypeOf(zero)
	if t == nil {
		return keyMap
	}
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return keyMap
	}
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if jsonTag := f.Tag.Get("json"); jsonTag != "" {
			name := strings.Split(jsonTag, ",")[0]
			if name != "" && name != "-" {
				// store keyed by field name (f.Name) with form value taken from json tag name
				keyMap[f.Name] = c.FormValue(name)
			}
		}
	}

	// convert map -> T
	if res, err := mapToStruct[T](keyMap); err == nil {
		return res
	}
	return keyMap
}

// mapToStruct converts a map[string]string into a struct of type T by setting exported fields.
// It supports basic kinds: string, ints, uints, bool, floats. Fields are matched by struct field name.
func mapToStruct[T any](m map[string]string) (T, error) {
	var out T
	v := reflect.ValueOf(&out).Elem()
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		sf := t.Field(i)
		fv := v.Field(i)

		// only settable/exported fields
		if !fv.CanSet() {
			continue
		}

		// The FormDataParser stores values keyed by field name.
		val, ok := m[sf.Name]
		if !ok {
			continue
		}
		if val == "" {
			// skip empty strings (optional)
			continue
		}

		switch fv.Kind() {
		case reflect.String:
			fv.SetString(val)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			iv, err := strconv.ParseInt(val, 10, fv.Type().Bits())
			if err != nil {
				return out, err
			}
			fv.SetInt(iv)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			uv, err := strconv.ParseUint(val, 10, fv.Type().Bits())
			if err != nil {
				return out, err
			}
			fv.SetUint(uv)
		case reflect.Bool:
			bv, err := strconv.ParseBool(val)
			if err != nil {
				return out, err
			}
			fv.SetBool(bv)
		case reflect.Float32, reflect.Float64:
			fv64, err := strconv.ParseFloat(val, fv.Type().Bits())
			if err != nil {
				return out, err
			}
			fv.SetFloat(fv64)
		// Add more kinds (slices, nested structs, time.Time etc.) as needed.
		default:
			// unsupported kind: ignore or return error as desired
			continue
		}
	}
	return out, nil
}

func FF[T any](c *fiber.Ctx) (T, error) {
	var out T
	var prod T
	prodValue := reflect.ValueOf(&prod).Elem() // Get a settable Value
	prodTag := reflect.TypeOf(out)
	// Loop through fields by index and assign values dynamically
	for i := 0; i < prodValue.NumField(); i++ {
		field := prodValue.Field(i)
		// fieldName := prodValue.Type().Field(i).Name
		tag := prodTag.Field(i).Tag
		if !field.CanSet() {
			continue
		}
		var jsonName string
		if jsonTag := tag.Get("json"); jsonTag == "" {
			jsonName = strings.Split(jsonTag, ",")[0]
			if jsonName != "" && jsonName != "-" {
				// store keyed by field name (f.Name) with form value taken from json tag name
				continue
			}
			continue
		}
		switch field.Kind() {
		case reflect.String:
			field.SetString(c.FormValue(jsonName))
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			iv, err := strconv.ParseInt(c.FormValue(jsonName), 10, field.Type().Bits())
			if err != nil {
				return out, err
			}
			field.SetInt(iv)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			uv, err := strconv.ParseUint(c.FormValue(jsonName), 10, field.Type().Bits())
			if err != nil {
				return out, err
			}
			field.SetUint(uv)
		case reflect.Bool:
			bv, err := strconv.ParseBool(c.FormValue(jsonName))
			if err != nil {
				return out, err
			}
			field.SetBool(bv)
		case reflect.Float32, reflect.Float64:
			field64, err := strconv.ParseFloat(c.FormValue(jsonName), field.Type().Bits())
			if err != nil {
				return out, err
			}
			field.SetFloat(field64)
		// Add more kinds (slices, nested structs, time.Time etc.) as needed.
		default:
			// unsupported kind: ignore or return error as desired
			continue
		}
	}
	return prod, nil
}
