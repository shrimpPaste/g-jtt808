package binary

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"io/ioutil"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// Writer JTT808数据缓冲写入对象
type Writer struct {
	*bytes.Buffer
}

// WriteBCD 写入BCD编码
func (w Writer) WriteBCD(s string, length int) error {
	if len(s)%4 != 0 {
		for i := 0; i < len(s)%4; i++ {
			s = "0" + s
		}
	}

	b, err := hex.DecodeString(s)
	if err != nil {
		return err
	}

	var data []byte

	if length != 0 {
		for i := len(b); i < length; i++ {
			data = append(data, 0x00)
		}
		data = append(data, b...)
	}

	_, err = w.Write(data)
	return err
}

func (w Writer) WriteString(str string, maxLength int, from StartingEnd) error {
	reader := bytes.NewReader([]byte(str))
	data, err := ioutil.ReadAll(transform.NewReader(reader, simplifiedchinese.GB18030.NewEncoder()))
	if err != nil {
		return err
	}

	err = w.WriteData(data, maxLength, from)
	return err
}

// WriteUint32 写入Uint32数据
func (w Writer) WriteUint32(data uint32) error {
	var temp = make([]byte, 4)

	binary.BigEndian.PutUint32(temp[:], data)

	_, err := w.Write(temp[:])
	return err
}

// WriteUint16 写入Uint16数据
func (w Writer) WriteUint16(data uint16) error {
	var temp = make([]byte, 2)

	binary.BigEndian.PutUint16(temp[:], data)

	_, err := w.Write(temp[:])
	return err
}

// WriteUint8 写入Uint8数据
func (w Writer) WriteUint8(data uint8) error {
	_, err := w.Write([]byte{data})
	return err
}

// WriteData 写入数据
func (w Writer) WriteData(data []byte, maxLength int, from StartingEnd) error {
	ew := ErrWrite{W: w}
	difference := maxLength - len(data)

	if from == Front {
		ew.writePlaceholder(difference)
		ew.Write(data)

	} else {
		ew.Write(data)
		ew.writePlaceholder(difference)
	}
	return ew.Err
}

func (w Writer) writePlaceholder(difference int) error {
	for i := 0; i < difference; i++ {
		if err := w.WriteByte(0x00); err != nil {
			return err
		}
	}
	return nil
}

// NewWriter 创建新的JTT808数据缓冲写入对象
func NewWriter() Writer {
	return Writer{
		Buffer: &bytes.Buffer{},
	}
}

// StartingEnd 起始端
type StartingEnd int

const (
	// Front 前置占位
	Front StartingEnd = iota

	// Behind 后置占位
	Behind
)

type ErrWrite struct {
	W   Writer
	Err error
}

func (ew *ErrWrite) Write(p []byte) {
	if ew.Err != nil {
		return
	}
	_, ew.Err = ew.W.Write(p)
}

func (ew ErrWrite) WriteByte(c byte) {
	if ew.Err != nil {
		return
	}
	ew.Err = ew.W.WriteByte(c)
}

func (ew ErrWrite) WriteUint8(data uint8) {
	if ew.Err != nil {
		return
	}
	ew.Err = ew.W.WriteUint8(data)
}

func (ew ErrWrite) WriteUint16(data uint16) {
	if ew.Err != nil {
		return
	}
	ew.Err = ew.W.WriteUint16(data)
}

func (ew ErrWrite) WriteUint32(data uint32) {
	if ew.Err != nil {
		return
	}
	ew.Err = ew.W.WriteUint32(data)
}

func (ew ErrWrite) WriteBCD(s string, length int) {
	if ew.Err != nil {
		return
	}
	ew.Err = ew.W.WriteBCD(s, length)
}

func (ew ErrWrite) WriteString(s string, maxLength int, from StartingEnd) {
	if ew.Err != nil {
		return
	}
	ew.Err = ew.W.WriteString(s, maxLength, from)
}

func (ew ErrWrite) writePlaceholder(d int) {
	if ew.Err != nil {
		return
	}
	ew.Err = ew.W.writePlaceholder(d)
}
