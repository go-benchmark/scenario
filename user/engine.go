package user

type engineType string

type engine struct {
	id string
}

const (
	softSecurity engineType = "softSecurity"
	location     engineType = "location"
)
const (
	// CmdStart cmd start command
	CmdStart = "start"
	// CmdStop cmd Stop command
	CmdStop = "stop"
	// DeviceMo device with type is MO
	DeviceMo = "mo"
	// DeviceBot device with type is bot
	DeviceBot = "bot"
)
