package bb

import (
	"encoding/binary"
	"math"
	"reflect"
)

func serializeInt(field reflect.Value, buf []byte) []byte {
	size := field.Type().Size()
	data := make([]byte, size)
	switch size {
	case 2:
		binary.BigEndian.PutUint16(data, uint16(field.Int()))
	case 4:
		binary.BigEndian.PutUint32(data, uint32(field.Int()))
	case 8:
		binary.BigEndian.PutUint64(data, uint64(field.Int()))
	}
	return append(buf, data...)
}

func serializeUint(field reflect.Value, buf []byte) []byte {
	size := field.Type().Size()
	data := make([]byte, size)
	switch size {
	case 2:
		binary.BigEndian.PutUint16(data, uint16(field.Uint()))
	case 4:
		binary.BigEndian.PutUint32(data, uint32(field.Uint()))
	case 8:
		binary.BigEndian.PutUint64(data, uint64(field.Uint()))
	}
	return append(buf, data...)
}

func serializeFloat(field reflect.Value, buf []byte) []byte {
	size := field.Type().Size()
	data := make([]byte, size)
	switch field.Kind() {
	case reflect.Float32:
		binary.BigEndian.PutUint32(data, math.Float32bits(float32(field.Float())))
	case reflect.Float64:
		binary.BigEndian.PutUint64(data, math.Float64bits(field.Float()))
	}
	return append(buf, data...)
}

func serializeString(field reflect.Value, buf []byte) []byte {
	str := field.String()
	strLen := uint16(len(str))
	lenData := make([]byte, 2)
	binary.BigEndian.PutUint16(lenData, strLen)
	buf = append(buf, lenData...)
	buf = append(buf, []byte(str)...)
	return buf
}

func serializeBool(field reflect.Value, buf []byte) []byte {
	if field.Bool() {
		buf = append(buf, byte(1))
	} else {
		buf = append(buf, byte(0))
	}
	return buf
}
