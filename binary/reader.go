package binary

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"
	"time"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

type Reader struct {
	*bytes.Reader
}

type ErrReader struct {
	R   Reader
	Err error
}

// ReadBCD 读取BCD编码
func (r *Reader) ReadBCD(length int) (string, error) {
	result := make([]byte, length)
	if _, err := r.Read(result); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", result), nil
}

// ReadUint16 读取双字节
func (r *Reader) ReadUint16() (uint16, error) {
	temp := make([]byte, 2)

	n, err := r.Read(temp[:])

	if err != nil {
		return 0, err
	}

	if n != len(temp) {
		return 0, io.ErrUnexpectedEOF
	}

	return binary.BigEndian.Uint16(temp[:]), nil
}

// ReadUint32 读取四字节
func (r *Reader) ReadUint32() (uint32, error) {
	temp := make([]byte, 4)

	n, err := r.Read(temp[:])
	if err != nil {
		return 0, err
	}

	if n != len(temp) {
		return 0, io.ErrUnexpectedEOF
	}

	return binary.BigEndian.Uint32(temp[:]), nil
}

// ReadBytes 读取多个字节
func (r *Reader) ReadBytes(size int) ([]byte, error) {
	temp := make([]byte, size)

	if _, err := r.Read(temp); err != nil {
		return nil, err
	}

	return temp, nil
}

// ReaderString 读取字符串
func (r *Reader) ReaderString(size int) (string, error) {
	b, err := r.ReadBytes(size)
	if err != nil {
		return "", err
	}

	text, err := ioutil.ReadAll(transform.NewReader(NewReader(b), simplifiedchinese.GBK.NewDecoder()))
	if err != nil {
		return "", err
	}

	return string(text), nil
}

func NewReader(b []byte) Reader {
	return Reader{
		Reader: bytes.NewReader(b),
	}
}

func (er *ErrReader) ReadBCD(length int) string {
	var bcd string
	if er.Err != nil {
		return ""
	}
	bcd, er.Err = er.R.ReadBCD(length)
	return bcd
}

func (er *ErrReader) ReadBCDTime() time.Time {
	if er.R.Len() < 6 {
		er.Err = fmt.Errorf("时间bcd长度不足六位")
		return time.Time{}
	}

	var buf [6]byte
	n, err := er.R.Read(buf[:])
	if err != nil {
		er.Err = err
		return time.Time{}
	}

	if n != len(buf) {
		er.Err = io.ErrUnexpectedEOF
		return time.Time{}
	}
	return fromBCDTime(buf[:])
}

func (er *ErrReader) ReadUint16() uint16 {
	var data uint16
	if er.Err != nil {
		return 0
	}
	data, er.Err = er.R.ReadUint16()
	return data
}

func (er *ErrReader) ReadUint32() uint32 {
	var data uint32
	if er.Err != nil {
		return 0
	}
	data, er.Err = er.R.ReadUint32()
	return data
}

func (er *ErrReader) ReadBytes(size int) []byte {
	var data []byte
	if er.Err != nil {
		return []byte{0}
	}
	data, er.Err = er.R.ReadBytes(size)
	return data
}

func (er *ErrReader) ReadByte() byte {
	var b byte
	if er.Err != nil {
		return 0
	}
	b, er.Err = er.R.ReadByte()
	return b
}

func (er *ErrReader) ReadString(size int) string {
	var data string
	if er.Err != nil {
		return ""
	}
	data, er.Err = er.R.ReaderString(size)
	return data
}

// 转为time.Time
func fromBCDTime(bcd []byte) time.Time {
	if len(bcd) != 6 {
		fmt.Println("时间长度BCD不足六位")
		return time.Time{}
	}
	t, err := time.ParseInLocation(
		"20060102150405", "20"+bcdToString(bcd), time.Local)
	if err != nil {
		return time.Time{}
	}
	return t
}

func bcdToString(data []byte, ignorePadding ...bool) string {
	for {
		if len(data) == 0 {
			return ""
		}
		if data[0] != 0 {
			break
		}
		data = data[1:]
	}

	var b []byte
	for i := 0; i < len(data); i++ {
		b = append(b, data[i]&0xf0>>4+'0')
		b = append(b, data[i]&0x0f+'0')
	}

	if len(ignorePadding) == 0 || !ignorePadding[0] {
		for idx := range b {
			if b[idx] != '0' {
				return string(b[idx:])
			}
		}
	}
	return string(b)
}
