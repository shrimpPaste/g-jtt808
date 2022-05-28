package v2013

type LoadStatus string

const (
	NoLoad   LoadStatus = "00"
	HalfLoad            = "01"
	HoldLoad            = "10"
	FullLoad            = "11"
)

var LoadStatusMap = map[string]LoadStatus{
	"00": NoLoad,
	"01": HalfLoad,
	"10": HoldLoad,
	"11": FullLoad,
}
