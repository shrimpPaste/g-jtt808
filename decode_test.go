package jtt808

import (
	"testing"

	"github.com/mingkid/jtt808/message"
	bv2013 "github.com/mingkid/jtt808/message/body/v2013"
	hv2013 "github.com/mingkid/jtt808/message/head/v2013"
	"github.com/mingkid/jtt808/message/msgcomm"
)

func TestDecoder_M0100(t *testing.T) {
	var msg message.Message[hv2013.MsgHead, bv2013.M0100]
	d := NewDecoder(&msg)
	_ = d.Decode([]byte{126, 1, 0, 0, 39, 1, 48, 81, 25, 38, 117, 0, 128, 0, 44, 1, 44, 55, 48, 49, 49, 49, 66, 83, 74, 45, 65, 54, 66, 68, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 49, 49, 57, 50, 54, 55, 53, 0, 0, 0, 131, 126})
	if msg.Head.ID() != msgcomm.TermRegister {
		t.Fatalf("消息包ID解析错误，应为%d，实际为%d", msgcomm.TermRegister, msg.Head.ID())
	}
	if msg.Head.SerialNum() != 128 {
		t.Fatalf("消息包流水号解析错误，应为%d，实际为%d", 128, msg.Head.SerialNum())
	}
	if msg.Head.Phone() != "013051192675" {
		t.Fatalf("消息包手机号解析错误，应为%s，实际为%s", "013051192675", msg.Head.Phone())
	}
	if msg.Head.BodyProperty.IsSub() != false {
		t.Fatalf("消息包分包解析错误，应为%t，实际为%t", false, msg.Head.BodyProperty.IsSub())
	}
	if msg.Head.BodyProperty.BodyLength() != 39 {
		t.Fatalf("消息包消息长度解析错误，应为%d，实际为%d", 39, msg.Head.BodyProperty.BodyLength())
	}
	if msg.Head.BodyProperty.EncryptType() != msgcomm.None {
		t.Fatalf("消息包加密解析错误，应为%d，实际为%d", 39, msg.Head.BodyProperty.EncryptType())
	}
	if msg.Body.ProvincialID() != 44 {
		t.Fatalf("消息包省域ID解析错误，应为%d，实际为%d", 44, msg.Body.ProvincialID())
	}
	if msg.Body.CityID() != 300 {
		t.Fatalf("消息包市县域ID解析错误，应为%d，实际为%d", 300, msg.Body.CityID())
	}
	if msg.Body.ManufacturerID() != "70111" {
		t.Fatalf("消息包制造商ID解析错误，应为%s，实际为%s", "70111", msg.Body.ManufacturerID())
	}
	if msg.Body.TermModel() != "BSJ-A6BD" {
		t.Fatalf("消息包解析错误，应为%s，实际为%s", "BSJ-A6BD", msg.Body.TermModel())
	}
	if msg.Body.TermID() != "1192675" {
		t.Fatalf("消息包终端ID解析错误，应为%s，实际为%s", "1192675", msg.Body.TermID())
	}
	if msg.Body.LicensePlateColor() != 0 {
		t.Fatalf("消息包车牌颜色解析错误，应为%d，实际为%d", 0, msg.Body.LicensePlateColor())
	}
	if msg.Body.LicensePlate() != "" {
		t.Fatalf("消息包车牌标识解析错误，应为%s，实际为%s", "", msg.Body.LicensePlate())
	}
}

func TestDecoder_M0102(t *testing.T) {
	var msg message.Message[hv2013.MsgHead, bv2013.M0102]
	d := NewDecoder(&msg)
	_ = d.Decode([]byte{126, 1, 2, 0, 32, 1, 48, 81, 25, 99, 56, 0, 106, 55, 54, 50, 54, 102, 53, 49, 57, 56, 48, 53, 51, 102, 98, 50, 99, 48, 49, 100, 100, 48, 101, 98, 101, 97, 100, 101, 54, 48, 99, 102, 51, 109, 126})
	if msg.Head.ID() != msgcomm.TermAuth {
		t.Fatalf("消息包ID解析错误，应为%d，实际为%d", msgcomm.TermAuth, msg.Head.ID())
	}
	if msg.Head.SerialNum() != 106 {
		t.Fatalf("消息包流水号解析错误，应为%d，实际为%d", 106, msg.Head.SerialNum())
	}
	if msg.Head.Phone() != "013051196338" {
		t.Fatalf("消息包手机号解析错误，应为%s，实际为%s", "013051196338", msg.Head.Phone())
	}
	if msg.Head.BodyProperty.IsSub() != false {
		t.Fatalf("消息包分包解析错误，应为%t，实际为%t", false, msg.Head.BodyProperty.IsSub())
	}
	if msg.Head.BodyProperty.BodyLength() != 32 {
		t.Fatalf("消息包消息长度解析错误，应为%d，实际为%d", 32, msg.Head.BodyProperty.BodyLength())
	}
	if msg.Head.BodyProperty.EncryptType() != msgcomm.None {
		t.Fatalf("消息包加密解析错误，应为%d，实际为%d", 39, msg.Head.BodyProperty.EncryptType())
	}
	if msg.Body.Token() != "7626f5198053fb2c01dd0ebeade60cf3" {
		t.Fatalf("消息包鉴权码解析错误，应为%s，实际为%s", "7626f5198053fb2c01dd0ebeade60cf3", msg.Body.Token())
	}
}
