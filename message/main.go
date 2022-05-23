package message

// Message 消息模型
type Message[THead, TBody any] struct {
	Head THead
	Body TBody
}

func CalcChecksum(b []byte) (result byte) {
	result = b[0] // 取第0个，从第1个开始异或
	for i := 1; i < len(b); i++ {
		result = result ^ b[i]
	}
	return
}
