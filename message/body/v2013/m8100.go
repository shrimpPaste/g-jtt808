package v2013

type M8100 struct {
	RawSerialNumber uint16 `jtt808:""`
	RawResult       uint8  `jtt808:""`
	RawToken        string `jtt808:"0,string"`
}

func (m M8100) SerialNumber() uint16 {
	return m.RawSerialNumber
}

func (m *M8100) SetSerialNumber(num uint16) *M8100 {
	m.RawSerialNumber = num
	return m
}

func (m M8100) Result() M8100Result {
	return M8100Result(m.RawResult)
}

func (m *M8100) SetResult(num M8100Result) *M8100 {
	m.RawResult = uint8(num)
	return m
}

func (m M8100) Token() string {
	return m.RawToken
}

func (m *M8100) SetToken(t string) *M8100 {
	m.RawToken = t
	return m
}

type M8100Result uint8

const (
	// Success 成功
	Success M8100Result = iota

	// CarRegistered 车辆已被注册
	CarRegistered

	// CarNotInDB 数据库中无该车辆
	CarNotInDB

	// TermRegistered 终端已被注册
	TermRegistered

	// TermNotInDB 数据库中无该终端
	TermNotInDB
)
