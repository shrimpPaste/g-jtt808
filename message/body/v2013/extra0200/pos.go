package extra0200

// PositionType 位置类型
type PositionType byte

const (
	PositionTypeForNone PositionType = iota
	// PositionTypeForRound 圆形区域
	PositionTypeForRound
	// PositionTypeForRectangle 矩形区域
	PositionTypeForRectangle
	// PositionTypeForPolygon 多边形区域
	PositionTypeForPolygon
	// PositionTypeForRoad 线路
	PositionTypeForRoad
)

// IOType 进出类型
type IOType byte

const (
	// M0200IOTypeForIn 进
	M0200IOTypeForIn IOType = iota
	// M0200IOTypeForOut 出
	M0200IOTypeForOut
)
