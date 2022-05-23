package jtt808

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/mingkid/jtt808/message"
	bv2013 "github.com/mingkid/jtt808/message/body/v2013"
	hv2013 "github.com/mingkid/jtt808/message/head/v2013"
	"github.com/mingkid/jtt808/message/msgcomm"
)

func Test8001Message(t *testing.T) {
	var msg message.Message[hv2013.MsgHead, bv2013.M8001]

	// 不分包
	msg.Head.SetID(msgcomm.PlatformCommResp)
	msg.Head.SetPhone("13680179679")
	msg.Head.SetSerialNum(123)
	msg.Body.SetMsgID(msgcomm.TermLocationRepose).SetSerialNumber(321).SetResult(bv2013.M8001Fail)

	e := NewEncoder(&msg)
	b, _ := e.Encode(message.CalcChecksum)
	if r := hex.EncodeToString(b); r != "7e80010005013680179679007b0141020001f37e" {
		t.Fatalf("消息包解析错误，应为%s，实际为%s", "7e80010005013680179679007b0141020001f37e", r)
	}

	//	分包
	msg.Head.BodyProperty = 0
	msg.Head.BodyProperty = msg.Head.BodyProperty.SetIsSub().(hv2013.MsgBodyProperty)
	msg.Head.PackagePacking = &hv2013.MsgPackagePacking{}
	msg.Head.PackagePacking.SetTotal(1).SetSerialNum(1)

	e = NewEncoder(&msg)
	b, _ = e.Encode(message.CalcChecksum)
	if r := hex.EncodeToString(b); r != "7e80012005013680179679007b000100010141020001d37e" {
		t.Fatalf("消息包分包解析错误，应为%s，实际为%s", "7e80012005013680179679007b000100010141020001d37e", r)
	}
}

func convHEX(data []byte) {
	result := hex.EncodeToString(data)
	for i := 2; i <= len(data)*2; i += 2 {
		fmt.Printf("%s ", result[i-2:i])
	}
	fmt.Println("")
}
