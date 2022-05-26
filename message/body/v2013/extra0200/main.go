package extra0200

import (
	"encoding/binary"
)

type Extra map[byte][]byte

func (e Extra) Mileage() uint32 {
	return binary.BigEndian.Uint32(e[0x01])
}

func (e Extra) Oil() uint16 {
	return binary.BigEndian.Uint16(e[0x02])
}

func (e Extra) Speed() uint16 {
	return binary.BigEndian.Uint16(e[0x03])
}

func (e Extra) ConformWarnID() uint16 {
	return binary.BigEndian.Uint16(e[0x04])
}

func (e Extra) OverSpeedWarn() OverSpeedWarn {
	return e[0x11]
}

func (e Extra) IOPositionWarn() IOPositionWarn {
	return e[0x12]
}

func (e Extra) DriveTooShortWarn() DriveTooShortWarn {
	return e[0x13]
}

func (e Extra) VehSignalStatus() VehSignalStatus {
	return VehSignalStatus(binary.BigEndian.Uint32(e[0x25]))
}

func (e Extra) IOStatus() IOStatus {
	return IOStatus(binary.BigEndian.Uint16(e[0x2A]))
}

func (e Extra) Analog() uint32 {
	return binary.BigEndian.Uint32(e[0x2B])
}

func (e Extra) SignalStrength() byte {
	return e[0x30][0]
}

func (e Extra) GNSSQty() byte {
	return e[0x31][0]
}
