package enum

import (
	"reflect"
)

func New[T any]() (enum T) {
	enumValue := reflect.ValueOf(&enum).Elem()
	enumType := reflect.TypeOf(&enum).Elem()
	for i := 0; i < enumValue.NumField(); i++ {
		field := enumValue.Field(i)
		fieldType := enumType.Field(i)
		if !field.CanSet() {
			continue
		}
		switch k := field.Kind(); k {
		default:
			panic("enum field must be string or integer")
		case reflect.String:
			field.SetString(fieldType.Name)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			field.SetInt(int64(i))
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			field.SetUint(uint64(i))
		}
	}
	return enum
}
