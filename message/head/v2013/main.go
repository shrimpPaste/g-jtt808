package v2013

import "github.com/mingkid/jtt808/message/msgcomm"

type MsgHead struct {
	id             uint16 `jtt808:""`
	BodyProperty   MsgBodyProperty
	phone          string `jtt808:"6,bcd"`
	serialNum      uint16 `jtt808:""`
	PackagePacking *MsgPackagePacking
}

// ID 消息ID
func (h MsgHead) ID() msgcomm.MsgID {
	return msgcomm.MsgID(h.id)
}

// SetID 设置消息ID
func (h *MsgHead) SetID(id msgcomm.MsgID) {
	h.id = uint16(id)
}

// SerialNum 消息流水号
func (h MsgHead) SerialNum() uint16 {
	return h.serialNum
}

// SetSerialNum 设置消息流水号
func (h *MsgHead) SetSerialNum(num uint16) {
	h.serialNum = num
}

// Phone 终端手机号
func (h MsgHead) Phone() string {
	return h.phone
}

// SetPhone 设置终端手机号
func (h *MsgHead) SetPhone(p string) {
	h.phone = p
}
