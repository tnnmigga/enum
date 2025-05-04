package enum

import (
	"fmt"
	"reflect"
	"strconv"
)

// Package enum provides a generic mechanism to initialize enumeration-like structs in Go.
// It uses reflection to populate struct fields based on their names (for strings),
// indices (for integers), or custom values specified in "enum" tags.
// Supports string, integer (signed or unsigned), and nested struct fields.
// Nested structs are initialized recursively. Pointer fields are not supported.
// Panics on errors, such as non-struct types, unsupported field types, invalid tags,
// or integer overflows.
//
// Example usage:
//  HttpStatus := enum.New[struct {
//      Code struct {
//          StatusOK                  int `enum:"200"`
//          StatusNotFound            int `enum:"404"`
//          StatusInternalServerError int `enum:"500"`
//      }
//      Type struct {
//          StatusOK                  string
//          StatusNotFound            string
//          StatusInternalServerError string
//      }
//  }]()
//
//  fmt.Println(HttpStatus.Code.StatusOK)                  // Outputs: 200
//  fmt.Println(HttpStatus.Code.StatusNotFound)            // Outputs: 404
//  fmt.Println(HttpStatus.Code.StatusInternalServerError) // Outputs: 500
//  fmt.Println(HttpStatus.Type.StatusOK)                  // Outputs: StatusOK
//  fmt.Println(HttpStatus.Type.StatusNotFound)            // Outputs: StatusNotFound
//  fmt.Println(HttpStatus.Type.StatusInternalServerError) // Outputs: StatusInternalServerError

// New initializes an enum instance of type T, which must be a struct.
// Fields are populated based on their names (for strings), indices (for integers),
// or values specified in the "enum" tag. Supports nested structs, which are initialized
// recursively. Pointer fields are not allowed. Panics if T is not a struct, if unsupported
// field types (including pointers) are used, or if integer values overflow the target field type.
func New[T any]() T {
	var enum T
	enumVal := reflect.ValueOf(&enum).Elem()
	enumType := reflect.TypeOf(&enum).Elem()

	// Initialize the struct recursively.
	initialize(enumVal, enumType)
	return enum
}

// initialize recursively initializes a struct, handling its fields and nested structs.
// Panics on errors, such as non-struct types, unsupported field types, invalid tags,
// or integer overflows.
func initialize(val reflect.Value, typ reflect.Type) {
	// Ensure the type is a struct.
	if typ.Kind() != reflect.Struct {
		panic(fmt.Sprintf("type %s is not a struct", typ))
	}

	// Iterate over all fields of the struct.
	for i := 0; i < val.NumField(); i++ {
		fieldVal := val.Field(i)
		fieldType := typ.Field(i)

		// Skip unexported fields that cannot be set.
		if !fieldVal.CanSet() {
			continue
		}

		// Get the "enum" tag, if present.
		tagVal := fieldType.Tag.Get("enum")

		// Handle field based on its type.
		fieldKind := fieldType.Type.Kind()

		// Check for disallowed pointer types.
		if fieldKind == reflect.Ptr {
			panic(fmt.Sprintf("field %s: pointer types are not supported", fieldType.Name))
		}

		// Handle nested structs recursively.
		if fieldKind == reflect.Struct {
			initialize(fieldVal, fieldType.Type)
			continue
		}

		// Handle basic types (string or integer).
		switch fieldKind {
		case reflect.String:
			// Use field name as default value, or tag if provided.
			value := fieldType.Name
			if tagVal != "" {
				value = tagVal
			}
			fieldVal.SetString(value)

		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			// Use field index as default value, or parse tag if provided.
			value := int64(i)
			if tagVal != "" {
				parsedVal, err := strconv.ParseInt(tagVal, 10, 64)
				if err != nil {
					panic(fmt.Sprintf("field %s: invalid enum tag: %v", fieldType.Name, err))
				}
				value = parsedVal
			}
			// Check for integer overflow.
			if err := checkIntOverflow(value, fieldKind); err != nil {
				panic(fmt.Sprintf("field %s: %v", fieldType.Name, err))
			}
			fieldVal.SetInt(value)

		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			// Use field index as default value, or parse tag if provided.
			value := uint64(i)
			if tagVal != "" {
				parsedVal, err := strconv.ParseUint(tagVal, 10, 64)
				if err != nil {
					panic(fmt.Sprintf("field %s: invalid enum tag: %v", fieldType.Name, err))
				}
				value = parsedVal
			}
			// Check for unsigned integer overflow.
			if err := checkUintOverflow(value, fieldKind); err != nil {
				panic(fmt.Sprintf("field %s: %v", fieldType.Name, err))
			}
			fieldVal.SetUint(value)

		default:
			panic(fmt.Sprintf("field %s: unsupported type %s; only string, integer, or struct types are allowed", fieldType.Name, fieldKind))
		}
	}
}

// checkIntOverflow verifies if the value fits within the range of the specified signed integer type.
// Returns an error if the value overflows; used to trigger a panic in the caller.
func checkIntOverflow(value int64, kind reflect.Kind) error {
	switch kind {
	case reflect.Int8:
		if value < -1<<7 || value > 1<<7-1 {
			return fmt.Errorf("value %d overflows int8 range [-128, 127]", value)
		}
	case reflect.Int16:
		if value < -1<<15 || value > 1<<15-1 {
			return fmt.Errorf("value %d overflows int16 range [-32768, 32767]", value)
		}
	case reflect.Int32:
		if value < -1<<31 || value > 1<<31-1 {
			return fmt.Errorf("value %d overflows int32 range [-2147483648, 2147483647]", value)
		}
	case reflect.Int, reflect.Int64:
		// No additional check needed for int/int64, as value is already int64.
	}
	return nil
}

// checkUintOverflow verifies if the value fits within the range of the specified unsigned integer type.
// Returns an error if the value overflows; used to trigger a panic in the caller.
func checkUintOverflow(value uint64, kind reflect.Kind) error {
	switch kind {
	case reflect.Uint8:
		if value > 1<<8-1 {
			return fmt.Errorf("value %d overflows uint8 range [0, 255]", value)
		}
	case reflect.Uint16:
		if value > 1<<16-1 {
			return fmt.Errorf("value %d overflows uint16 range [0, 65535]", value)
		}
	case reflect.Uint32:
		if value > 1<<32-1 {
			return fmt.Errorf("value %d overflows uint32 range [0, 4294967295]", value)
		}
	case reflect.Uint, reflect.Uint64:
		// No additional check needed for uint/uint64, as value is already uint64.
	}
	return nil
}
