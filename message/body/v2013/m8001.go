package v2013

type M8001 struct {
	serialNumber uint16 `jtt808:""`
	msgID        uint16 `jtt808:""`
	result       uint8  `jtt808:""`
}

func (m M8001) SerialNumber() uint16 {
	return m.serialNumber
}

func (m M8001) MsgID() uint16 {
	return m.msgID
}

func (m M8001) Result() M8001Result {
	return M8001Result(m.result)
}

func (m *M8001) SetSerialNumber(number uint16) *M8001 {
	m.serialNumber = number
	return m
}

func (m *M8001) SetMsgID(id uint16) *M8001 {
	m.msgID = id
	return m
}

func (m *M8001) SetResult(result M8001Result) *M8001 {
	m.result = uint8(result)
	return m
}

type M8001Result uint8

const (
	// M8001Success 成功
	M8001Success M8001Result = iota

	// M8001Fail 失败
	M8001Fail

	// M8001MsgErr 消息有误
	M8001MsgErr

	// M8001Unsupported 不支持
	M8001Unsupported

	// M8001WarnConfirm 报警处理确认
	M8001WarnConfirm
)
