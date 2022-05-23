package v2013

import "testing"

func TestMsgPackagePacking_SerialNum(t *testing.T) {
	var p MsgPackagePacking
	p.SetSerialNum(123)
	if p.SerialNum() != 123 {
		t.Fatalf("设置包序号错误，应为%d，实际为%d", 123, p.SerialNum())
	}
}

func TestMsgPackagePacking_Total(t *testing.T) {
	var p MsgPackagePacking
	p.SetTotal(2)
	if p.Total() != 2 {
		t.Fatalf("设置消息总包数错误，应为%d，实际为%d", 2, p.Total())
	}
}
