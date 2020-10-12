package config

// Options main options
type Options struct {
	Host        string
	MQTTPrefix  string
	VirtualUser int
	UC          UserConfiguration
	DC          DeviceConfiguration
	SC          ServiceConfiguration
}

// UserConfiguration represent to configuration of users
type UserConfiguration struct {
	RunServiceDelay    float64
	GetDataInterval    float64
	RealtimeInterval   float64
	RealtimePeriod     float64
	RealtimeHBInterval int
	StartServiceDelay  float64
}

// DeviceConfiguration represent to configuration of devices
type DeviceConfiguration struct {
	RealtimeInterval float64
	HistoryInterval  float64
}

// ServiceConfiguration represent to configuration of services
type ServiceConfiguration struct {
	RealtimeLength int
}
