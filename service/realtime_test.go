package service

import (
	"context"
	"testing"
)

func TestService_RealtimeHandler(t *testing.T) {
	type args struct {
		ctx              context.Context
		realtimeInterval float64
		pub              func(context.Context, string, byte, []byte) error
	}
	tests := []struct {
		name    string
		s       *Service
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.RealtimeHandler(tt.args.ctx, tt.args.realtimeInterval, tt.args.pub); (err != nil) != tt.wantErr {
				t.Errorf("Service.RealtimeHandler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestService_RealtimeConfig(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		s       *Service
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.RealtimeConfig(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Service.RealtimeConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
