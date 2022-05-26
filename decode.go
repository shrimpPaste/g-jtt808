package jtt808

import (
	"reflect"
	"strings"

	"github.com/mingkid/jtt808/binary"
	"github.com/mingkid/jtt808/util"
)

type Decoder struct {
	msg any
}

func (d Decoder) Decode(b []byte) (err error) {
	// 去除首位标识位，校验码
	b = binary.EliminateChecksum(binary.EliminateIdentityBit(b))
	er := binary.ErrReader{R: binary.NewReader(b)}

	v := reflect.ValueOf(d.msg).Elem()
	d.decode(v.FieldByName("Head"), er)
	d.decode(v.FieldByName("Body"), er)

	return er.Err
}

func (d Decoder) decode(v reflect.Value, r binary.ErrReader) {
	for i := 0; i < v.NumField(); i++ {
		fv := v.Field(i)

		switch fv.Kind() {
		case reflect.Uint8:
			fv.SetUint(uint64(r.ReadByte()))
		case reflect.Uint16:
			fv.SetUint(uint64(r.ReadUint16()))
		case reflect.Uint32:
			fv.SetUint(uint64(r.ReadUint32()))
		case reflect.String:
			d.decodeString(r, fv, v.Type().Field(i))
		case reflect.Slice:
			d.decodeSlice(r, fv, v.Type().Field(i))
		case reflect.Struct:
			d.decode(fv, r)
		case reflect.Ptr:
			if fv.IsNil() {
				break
			}
			d.decode(fv.Elem(), r)
		}
	}
}

func (d Decoder) decodeString(r binary.ErrReader, v reflect.Value, f reflect.StructField) {
	var (
		t      string
		length int
	)
	t, length, r.Err = util.Tag(f)

	switch t {
	case "bcd":
		v.Set(reflect.ValueOf(r.ReadBCD(length)))
	case "string":
		v.Set(reflect.ValueOf(strings.Trim(r.ReadString(r.R.Len()), "\u0000")))
	}
}

func (d Decoder) decodeSlice(r binary.ErrReader, v reflect.Value, f reflect.StructField) {
	var length int
	if _, length, r.Err = util.Tag(f); length == 0 {
		v.Set(reflect.ValueOf(r.ReadBytes(r.R.Len())))
	} else {
		v.Set(reflect.ValueOf(r.ReadBytes(length)))
	}
}

func NewDecoder(msg any) Decoder {
	return Decoder{msg: msg}
}
