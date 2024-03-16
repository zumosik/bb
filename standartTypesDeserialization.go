package bb

import (
	"encoding/binary"
	"errors"
	"math"
	"reflect"
)

var (
	ErrNotEnoughData  = errors.New("not enough data")
	ErrUnsuportedSize = errors.New("unsupported size")
)

func deserializeInt(data []byte, field reflect.Value) ([]byte, error) {
	size := field.Type().Size()
	switch size {
	case 2:
		if len(data) < 2 {
			return nil, ErrNotEnoughData
		}
		intValue := int64(binary.BigEndian.Uint16(data))
		field.SetInt(intValue)
		return data[2:], nil
	case 4:
		if len(data) < 4 {
			return nil, ErrNotEnoughData
		}
		intValue := int64(binary.BigEndian.Uint32(data))
		field.SetInt(intValue)
		return data[4:], nil
	case 8:
		if len(data) < 8 {
			return nil, ErrNotEnoughData
		}
		intValue := int64(binary.BigEndian.Uint64(data))
		field.SetInt(intValue)
		return data[8:], nil
	}
	return nil, ErrUnsuportedSize
}

func deserializeUint(data []byte, field reflect.Value) ([]byte, error) {
	size := field.Type().Size()
	switch size {
	case 2:
		if len(data) < 2 {
			return nil, ErrNotEnoughData
		}
		uintValue := uint64(binary.BigEndian.Uint16(data))
		field.SetUint(uintValue)
		return data[2:], nil
	case 4:
		if len(data) < 4 {
			return nil, ErrNotEnoughData
		}
		uintValue := uint64(binary.BigEndian.Uint32(data))
		field.SetUint(uintValue)
		return data[4:], nil
	case 8:
		if len(data) < 8 {
			return nil, ErrNotEnoughData
		}
		uintValue := binary.BigEndian.Uint64(data)
		field.SetUint(uintValue)
		return data[8:], nil
	}
	return nil, ErrUnsuportedSize
}

func deserializeFloat(data []byte, field reflect.Value) ([]byte, error) {
	size := field.Type().Size()
	switch size {
	case 4:
		if len(data) < 4 {
			return nil, ErrNotEnoughData
		}
		floatValue := math.Float32frombits(binary.BigEndian.Uint32(data))
		field.SetFloat(float64(floatValue))
		return data[4:], nil
	case 8:
		if len(data) < 8 {
			return nil, ErrNotEnoughData
		}
		floatValue := math.Float64frombits(binary.BigEndian.Uint64(data))
		field.SetFloat(floatValue)
		return data[8:], nil
	}
	return nil, ErrUnsuportedSize
}

func deserializeString(data []byte, field reflect.Value) ([]byte, error) {
	if len(data) < 2 {
		return nil, ErrNotEnoughData
	}
	strLen := int(binary.BigEndian.Uint16(data))
	data = data[2:]

	if len(data) < strLen {
		return nil, ErrNotEnoughData
	}

	str := string(data[:strLen])
	field.SetString(str)
	return data[strLen:], nil
}

func deserializeBool(data []byte, field reflect.Value) []byte {
	//if len(data) < 1 {
	//	return nil, ErrNotEnoughData
	//} i think this have no reason

	if data[0] == 1 {
		field.SetBool(true)
	} else {

		field.SetBool(true)
	}

	return data[1:]
}
