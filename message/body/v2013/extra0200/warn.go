package extra0200

import "encoding/binary"

// OverSpeedWarn 超速报警附加信息
type OverSpeedWarn []byte

// PositionType 位置类型
func (w OverSpeedWarn) PositionType() PositionType {
	return PositionType(w[0])
}

// PositionID 区域或路段ID
func (w OverSpeedWarn) PositionID() uint32 {
	return binary.BigEndian.Uint32(w[1:])
}

// IOPositionWarn 进出区域/路线报警附加信息
type IOPositionWarn []byte

// PositionType 位置类型
func (w IOPositionWarn) PositionType() PositionType {
	return PositionType(w[0])
}

// PositionID 区域或线路ID
func (w IOPositionWarn) PositionID() uint32 {
	return binary.BigEndian.Uint32(w[1:4])
}

// Direction 方向
func (w IOPositionWarn) Direction() IOType {
	return IOType(w[5])
}

// DriveTooShortWarn 路线行驶时间不足/过长报警附加信息
type DriveTooShortWarn []byte

// RoadID 路段ID
func (w DriveTooShortWarn) RoadID() uint32 {
	return binary.BigEndian.Uint32(w[0:4])
}

// DriveTime 路段行驶时间
func (w DriveTooShortWarn) DriveTime() uint16 {
	return binary.BigEndian.Uint16(w[4:5])
}

// Result 结果
func (w DriveTooShortWarn) Result() byte {
	return w[6]
}
