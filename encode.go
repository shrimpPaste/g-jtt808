package jtt808

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/mingkid/jtt808/message/head"

	"github.com/mingkid/jtt808/binary"
)

type Encoder struct {
	message             any
	OnEncodedBodyHandle func(bodyBytes *[]byte)
}

func (e Encoder) Encode(calcChecksumHandle func([]byte) byte) ([]byte, error) {
	ew := binary.ErrWrite{W: binary.NewWriter()}
	v := reflect.ValueOf(e.message).Elem()

	var (
		hb []byte
		bb []byte
		bv = v.FieldByName("Body")
		hv = v.FieldByName("Head")
	)
	bb = e.encode(bv)
	hb, err := e.encodeHead(hv, uint16(len(bb)))
	if err != nil {
		return nil, err
	}

	ew.WriteByte(binary.IdentityBit)
	ew.Write(hb)
	ew.Write(bb)
	ew.WriteByte(calcChecksumHandle(ew.W.Bytes()[1:]))
	ew.WriteByte(binary.IdentityBit)
	if ew.Err != nil {
		return nil, ew.Err
	}

	return ew.W.Bytes(), ew.Err
}

func (e Encoder) encodeHead(v reflect.Value, bodyLen uint16) ([]byte, error) {
	pv := v.FieldByName("BodyProperty")
	p, err := pv.Interface().(head.BodyProperty).SetBodyLength(bodyLen)
	if err != nil {
		return nil, err
	}
	pv.Set(reflect.ValueOf(p))

	return e.encode(v), nil
}

func (e Encoder) encode(v reflect.Value) []byte {
	ew := binary.ErrWrite{W: binary.NewWriter()}

	for i := 0; i < v.NumField(); i++ {
		fv := v.Field(i)

		switch fv.Kind() {
		case reflect.Uint16:
			ew.WriteUint16(uint16(fv.Uint()))
		case reflect.Uint8:
			ew.WriteUint8(uint8(fv.Uint()))
		case reflect.String:
			e.encodeString(fv.String(), ew, v.Type().Field(i))
		case reflect.Struct:
			e.encode(fv)
		case reflect.Ptr:
			if fv.IsNil() {
				break
			}
			ew.Write(e.encode(fv.Elem()))
		}
	}

	return ew.W.Bytes()
}

func (e Encoder) encodeString(s string, w binary.ErrWrite, f reflect.StructField) {
	// 获取消息struct中的Tag
	tag, ok := f.Tag.Lookup("jtt808")
	if !ok {
		w.Err = fmt.Errorf("未定义数据长度")
	}

	var (
		tagItems    = strings.Split(tag, ",")
		t           = tagItems[1]
		length, err = strconv.Atoi(tagItems[0])
	)
	if err != nil {
		w.Err = fmt.Errorf("数据长度不正确，请按照：'长度,类型'格式来设置Tag")
	}

	// 根据Tag中的定义，为string类型的数据做不同的bytes转换
	switch t {
	case "bcd":
		w.WriteBCD(s, length)
	default:
		w.Err = fmt.Errorf("数据类型不正确，字符串只支持STRING|BCD|BYTE，并按照'长度,类型'格式来设置Tag，如：10,BCD")
	}
}

func NewEncoder(msg any) Encoder {
	return Encoder{
		message: msg,
		OnEncodedBodyHandle: func(bodyBytes *[]byte) {
			return
		},
	}
}
