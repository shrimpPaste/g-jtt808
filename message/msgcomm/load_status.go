package msgcomm

type LoadStatus string

const (
	NoLoad   LoadStatus = "00"
	HalfLoad            = "01"
	HoldLoad            = "10"
	FullLoad            = "11"
	IsExist             = "载客状态异常,请检查"
)

var loadStatusMap = map[string]LoadStatus{
	"00": NoLoad,
	"01": HalfLoad,
	"10": HoldLoad,
	"11": FullLoad,
}

func GetStatus(status string) LoadStatus {
	if _, ok := loadStatusMap[status]; !ok {
		return IsExist
	}
	return loadStatusMap[status]
}
