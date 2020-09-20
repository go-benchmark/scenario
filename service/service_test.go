package service

import (
	"reflect"
	"testing"
)

var s *Service

func init() {
	s = &Service{}
}
func TestNewService(t *testing.T) {
	type args struct {
		s              *Service
		realtimeLength int
	}
	tests := []struct {
		name string
		args args
		want *Service
	}{
		{
			name: "it should create a new service",
			args: args{
				s:              s,
				realtimeLength: 1,
			},
			want: &Service{
				OI: OperationInfo{
					RealtimeLength: 1,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewService(tt.args.s, tt.args.realtimeLength)
			tt.want.heartbeat = got.heartbeat

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
