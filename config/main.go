package config

import (
	"fmt"

	"github.com/spf13/viper"
)

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

// ConfigureOptions main config initial
func ConfigureOptions(vu int) (*Options, error) {
	opts := &Options{
		Host:        viper.GetString("api.endpoint"),
		MQTTPrefix:  viper.GetString("mqtt.prefix"),
		VirtualUser: vu,
		UC: UserConfiguration{
			RunServiceDelay:    viper.GetFloat64("testing.user.runServiceDelay"),
			GetDataInterval:    viper.GetFloat64("testing.user.getDataInterval"),
			RealtimeInterval:   viper.GetFloat64("testing.user.realtime.interval"),
			RealtimePeriod:     viper.GetFloat64("testing.user.realtime.period"),
			RealtimeHBInterval: viper.GetInt("testing.user.realtime.heartbeatInterval"),
			StartServiceDelay:  viper.GetFloat64("testing.user.startServiceDelay"),
		},
		DC: DeviceConfiguration{
			RealtimeInterval: viper.GetFloat64("testing.device.realtimeInterval"),
			HistoryInterval:  viper.GetFloat64("testing.device.historyInterval"),
		},
		SC: ServiceConfiguration{
			RealtimeLength: viper.GetInt("testing.service.realtimeLength"),
		},
	}

	return opts, nil
}

func init() {
	// read config from yaml file
	viper.SetConfigName("default") // name of config file (without extension)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")                                                    // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./config")                                             // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("$HOME/go/src/github.com/go-benchmark/scenario/config") // path to look for the config file in
	err := viper.ReadInConfig()                                                 // Find and read the config file
	if err != nil {                                                             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %s ", err))
	}
}
