package binary

// IdentityBit 标识位
const IdentityBit uint8 = 0x7e

const Escape uint8 = 0x7d

// EliminateIdentityBit 去除首尾校验码
func EliminateIdentityBit(data []byte) []byte {
	if data[0] == IdentityBit {
		data = data[1:]
	}

	if data[len(data)-1] == IdentityBit {
		data = data[:len(data)-1]
	}

	return data
}

// EliminateChecksum 去除校验码
func EliminateChecksum(data []byte) []byte {
	return data[:len(data)-1]
}
