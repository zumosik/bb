package bb

import (
	"math"
	"reflect"
	"testing"
)

type allFieldsTest struct {
	I16  int16
	I32  int32
	I64  int64
	Ui16 uint16
	Ui32 uint32
	Ui64 uint64
	F32  float32
	F64  float64
	S1   string
	S2   string
}

func TestMarshal(t *testing.T) {
	type args struct {
		v any
	}
	tests := []struct {
		name    string
		args    args
		wantBuf []byte
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				v: struct {
					IntValue    int16
					UintValue   uint32
					FloatValue  float32
					StringValue string
				}{
					IntValue:    42,
					UintValue:   12345,
					FloatValue:  3.14,
					StringValue: "hello",
				},
			},
			wantBuf: []byte{
				0x00, 0x2a, // Int16: 42
				0x00, 0x00, 0x30, 0x39, // Uint32: 12345
				0x40, 0x48, 0xf5, 0xc3, // Float32: 3.14
				0x00, 0x05, // Length of string: 5
				0x68, 0x65, 0x6c, 0x6c, 0x6f, // String: "hello"
			},
			wantErr: false,
		},
		{
			name: "ok 2",
			args: args{
				struct {
					IntValue    int16
					UintValue   uint32
					FloatValue  float32
					StringValue string
				}{
					IntValue:    -100,
					UintValue:   54321,
					FloatValue:  float32(math.Inf(1)),
					StringValue: "world",
				},
			},
			wantBuf: []byte{
				0xff, 0x9c, // Int16: -100
				0x00, 0x00, 0xd4, 0x31, // Uint32: 54321
				0x7f, 0x80, 0x00, 0x00, // Float32: +Inf
				0x00, 0x05, // Length of string: 5
				0x77, 0x6f, 0x72, 0x6c, 0x64, // String: "world"
			},
			wantErr: false,
		},
		{
			name: "invalid data",
			args: args{
				v: struct {
					Invalid interface{}
				}{
					Invalid: struct {
						a string
						b int
					}{},
				},
			},
			wantBuf: nil,
			wantErr: true,
		},
		{
			name: "invalid type",
			args: args{
				v: "a",
			},
			wantBuf: nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBuf, err := Marshal(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("Marshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotBuf, tt.wantBuf) {
				t.Errorf("Marshal() gotBuf = %v, want %v", gotBuf, tt.wantBuf)
			}
		})
	}
}

func TestMarshalUnmarshal(t *testing.T) {
	tests := []struct {
		name string
		data allFieldsTest
	}{
		{
			name: "ok",
			data: allFieldsTest{
				I16:  -666,
				I32:  222,
				I64:  -87,
				Ui16: 5,
				Ui32: 999,
				Ui64: 222,
				F32:  2525.001,
				F64:  math.Inf(2),
				S1:   "hello",
				S2:   "world!",
			},
		},
		{
			name: "ok 2",
			data: allFieldsTest{
				I16:  -6666,
				I32:  22772,
				I64:  -8790,
				Ui16: 55555,
				Ui32: 999899,
				Ui64: 0,
				F32:  2525.255225,
				F64:  math.Inf(4),
				S1:   "some rand string",
				S2:   "lorem ipsum",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBuf, err := Marshal(tt.data)
			if err != nil {
				t.Fatalf("Error marshalling data: %v", err)
			}

			var unmarshalledData allFieldsTest
			err = Unmarshal(gotBuf, &unmarshalledData)
			if err != nil {
				t.Fatalf("Error unmarshalling data: %v", err)
			}

			if !reflect.DeepEqual(tt.data, unmarshalledData) {
				t.Errorf("Original and unmarshalled data are not equal, data before: %v, data after: %v", tt.data, unmarshalledData)
			}
		})
	}
}
