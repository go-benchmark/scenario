package config

import (
	"reflect"
	"testing"
)

func TestConfigureOptions(t *testing.T) {
	tests := []struct {
		name    string
		vu      int
		want    *Options
		wantErr bool
	}{
		{
			name: "it should load config success for 1 virtual user",
			vu:   1,
			want: &Options{
				Host:        "owi-fs-loadtest.veriksystems.com",
				MQTTPrefix:  "org/vz-loadtest",
				VirtualUser: 1,
				DC: DeviceConfiguration{
					RealtimeInterval: 1.0,
					HistoryInterval:  10.0,
				},
				SC: ServiceConfiguration{
					RealtimeLength: 6,
				},
				UC: UserConfiguration{
					RunServiceDelay:    10.0,
					StartServiceDelay:  900.0,
					GetDataInterval:    600.0,
					RealtimeInterval:   600.0,
					RealtimePeriod:     30.0,
					RealtimeHBInterval: 6,
				},
			},
		},
		{
			name: "it should load config success for 2 virtual user",
			vu:   2,
			want: &Options{
				Host:        "owi-fs-loadtest.veriksystems.com",
				MQTTPrefix:  "org/vz-loadtest",
				VirtualUser: 2,
				DC: DeviceConfiguration{
					RealtimeInterval: 1.0,
					HistoryInterval:  10.0,
				},
				SC: ServiceConfiguration{
					RealtimeLength: 6,
				},
				UC: UserConfiguration{
					RunServiceDelay:    10.0,
					StartServiceDelay:  900.0,
					GetDataInterval:    600.0,
					RealtimeInterval:   600.0,
					RealtimePeriod:     30.0,
					RealtimeHBInterval: 6,
				},
			},
		},
		{
			name: "it should load config success for 2000 virtual user",
			vu:   2000,
			want: &Options{
				Host:        "owi-fs-loadtest.veriksystems.com",
				MQTTPrefix:  "org/vz-loadtest",
				VirtualUser: 2000,
				DC: DeviceConfiguration{
					RealtimeInterval: 1.0,
					HistoryInterval:  10.0,
				},
				SC: ServiceConfiguration{
					RealtimeLength: 6,
				},
				UC: UserConfiguration{
					RunServiceDelay:    10.0,
					StartServiceDelay:  900.0,
					GetDataInterval:    600.0,
					RealtimeInterval:   600.0,
					RealtimePeriod:     30.0,
					RealtimeHBInterval: 6,
				},
			},
		},
		{
			name: "it should load config success for 10000 virtual user",
			vu:   10000,
			want: &Options{
				Host:        "owi-fs-loadtest.veriksystems.com",
				MQTTPrefix:  "org/vz-loadtest",
				VirtualUser: 10000,
				DC: DeviceConfiguration{
					RealtimeInterval: 1.0,
					HistoryInterval:  10.0,
				},
				SC: ServiceConfiguration{
					RealtimeLength: 6,
				},
				UC: UserConfiguration{
					RunServiceDelay:    10.0,
					StartServiceDelay:  900.0,
					GetDataInterval:    600.0,
					RealtimeInterval:   600.0,
					RealtimePeriod:     30.0,
					RealtimeHBInterval: 6,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConfigureOptions(tt.vu)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConfigureOptions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConfigureOptions() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
