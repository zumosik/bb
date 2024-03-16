package bb

import (
	"errors"
	"fmt"
	"reflect"
)

var (
	ErrInvalidData = errors.New("invalid data")
	ErrInvalidType = errors.New("invalid type")
)

// Marshal serializes struct, where fields can be int16, int32, int64, uint16, uint32, uint64, float32, float64, string, bool.
//
// Order of fields is important, also fields need to be exported
func Marshal(v interface{}) (buf []byte, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = ErrInvalidData
		}
	}()

	val := reflect.ValueOf(v)
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fmt.Println(field.Kind())

		if !field.CanInterface() {
			continue
		}

		switch field.Kind() {
		case reflect.Int16, reflect.Int32, reflect.Int64:
			buf = serializeInt(field, buf)
		case reflect.Uint16, reflect.Uint32, reflect.Uint64:
			buf = serializeUint(field, buf)
		case reflect.Float32, reflect.Float64:
			buf = serializeFloat(field, buf)
		case reflect.String:
			buf = serializeString(field, buf)
		case reflect.Bool:
			buf = serializeBool(field, buf)
		default:
			return nil, ErrInvalidType
		}
	}

	return buf, nil
}

// Unmarshal deserializes struct, where fields can be int16, int32, int64, uint16, uint32, uint64, float32, float64, string, bool.
//
// Order of fields is important
func Unmarshal(data []byte, v any) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = ErrInvalidData
		}
	}()

	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Ptr || val.IsNil() {
		return ErrInvalidType
	}

	val = val.Elem()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		if !field.CanInterface() {
			continue
		}

		switch field.Kind() {
		case reflect.Int16, reflect.Int32, reflect.Int64:
			data, err = deserializeInt(data, field)
		case reflect.Uint16, reflect.Uint32, reflect.Uint64:
			data, err = deserializeUint(data, field)
		case reflect.Float32, reflect.Float64:
			data, err = deserializeFloat(data, field)
		case reflect.String:
			data, err = deserializeString(data, field)
		case reflect.Bool:
			data = deserializeBool(data, field)
		default:
			return ErrInvalidType
		}

		if err != nil {
			return err
		}
	}

	return nil
}
