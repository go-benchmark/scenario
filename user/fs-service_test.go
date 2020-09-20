package user

import (
	"context"
	"reflect"
	"testing"
)

func TestUser_CreateService(t *testing.T) {
	type args struct {
		ctx context.Context
		ds  DeviceSet
		et  engineType
	}
	tests := []struct {
		name    string
		u       *User
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.u.CreateService(tt.args.ctx, tt.args.ds, tt.args.et); (err != nil) != tt.wantErr {
				t.Errorf("User.CreateService() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUser_AttachServices(t *testing.T) {
	type args struct {
		ctx        context.Context
		ds         DeviceSet
		serviceIds []string
	}
	tests := []struct {
		name    string
		u       *User
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.u.AttachServices(tt.args.ctx, tt.args.ds, tt.args.serviceIds); (err != nil) != tt.wantErr {
				t.Errorf("User.AttachServices() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUser_StartServices(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		u       *User
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.u.StartServices(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("User.StartServices() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUser_StopServices(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		u       *User
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.u.StopServices(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("User.StopServices() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUser_RunServices(t *testing.T) {
	type args struct {
		ctx context.Context
		cmd string
	}
	tests := []struct {
		name    string
		u       *User
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.u.RunServices(tt.args.ctx, tt.args.cmd); (err != nil) != tt.wantErr {
				t.Errorf("User.RunServices() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUser_RunService(t *testing.T) {
	type args struct {
		ctx  context.Context
		s    FSService
		cmd  string
		moID string
		bots []string
	}
	tests := []struct {
		name    string
		u       *User
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.u.RunService(tt.args.ctx, tt.args.s, tt.args.cmd, tt.args.moID, tt.args.bots); (err != nil) != tt.wantErr {
				t.Errorf("User.RunService() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUser_CheckConfig(t *testing.T) {
	type args struct {
		ctx context.Context
		s   FSService
	}
	tests := []struct {
		name    string
		u       *User
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.u.CheckConfig(tt.args.ctx, tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("User.CheckConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUser_GetConfig(t *testing.T) {
	type args struct {
		ctx context.Context
		s   FSService
	}
	tests := []struct {
		name    string
		u       *User
		args    args
		wantCs  ConfigService
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCs, err := tt.u.GetConfig(tt.args.ctx, tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("User.GetConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCs, tt.wantCs) {
				t.Errorf("User.GetConfig() = %v, want %v", gotCs, tt.wantCs)
			}
		})
	}
}

func TestUser_GetHistories(t *testing.T) {
	type args struct {
		ctx context.Context
		s   FSService
	}
	tests := []struct {
		name    string
		u       *User
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.u.GetHistories(tt.args.ctx, tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("User.GetHistories() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUser_CreateHeartbeat(t *testing.T) {
	type args struct {
		ctx context.Context
		s   FSService
	}
	tests := []struct {
		name    string
		u       *User
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.u.CreateHeartbeat(tt.args.ctx, tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("User.CreateHeartbeat() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUser_SetServiceParams(t *testing.T) {
	type args struct {
		ctx context.Context
		s   FSService
	}
	tests := []struct {
		name    string
		u       *User
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.u.SetServiceParams(tt.args.ctx, tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("User.SetServiceParams() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUser_GetServiceParams(t *testing.T) {
	type args struct {
		ctx context.Context
		s   FSService
	}
	tests := []struct {
		name    string
		u       *User
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.u.GetServiceParams(tt.args.ctx, tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("User.GetServiceParams() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
