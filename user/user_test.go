package user

import (
	"context"
	"reflect"
	"testing"

	"github.com/go-benchmark/scenario/device"
)

func TestUser_SignUp(t *testing.T) {
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
			if err := tt.u.SignUp(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("User.SignUp() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUser_Login(t *testing.T) {
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
			if err := tt.u.Login(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("User.Login() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUser_GetUserAccount(t *testing.T) {
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
			if err := tt.u.GetUserAccount(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("User.GetUserAccount() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUser_headers(t *testing.T) {
	tests := []struct {
		name string
		u    *User
		want map[string]string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.u.headers(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("User.headers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_CreateSsEngine(t *testing.T) {
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
			if err := tt.u.CreateSsEngine(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("User.CreateSsEngine() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUser_CreateLoEngine(t *testing.T) {
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
			if err := tt.u.CreateLoEngine(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("User.CreateLoEngine() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUser_createEngine(t *testing.T) {
	type args struct {
		ctx context.Context
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
			if err := tt.u.createEngine(tt.args.ctx, tt.args.et); (err != nil) != tt.wantErr {
				t.Errorf("User.createEngine() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUser_AttachDevice(t *testing.T) {
	type args struct {
		ctx context.Context
		d   *device.Device
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
			if err := tt.u.AttachDevice(tt.args.ctx, tt.args.d); (err != nil) != tt.wantErr {
				t.Errorf("User.AttachDevice() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUser_CreateDeviceSet(t *testing.T) {
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
			if err := tt.u.CreateDeviceSet(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("User.CreateDeviceSet() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUser_AddDevicesToDeviceSets(t *testing.T) {
	type args struct {
		ctx context.Context
		d   *device.Device
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
			if err := tt.u.AddDevicesToDeviceSets(tt.args.ctx, tt.args.d); (err != nil) != tt.wantErr {
				t.Errorf("User.AddDevicesToDeviceSets() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUser_AddServicesToDeviceSets(t *testing.T) {
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
			if err := tt.u.AddServicesToDeviceSets(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("User.AddServicesToDeviceSets() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUser_GetHistoriesByUser(t *testing.T) {
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
			if err := tt.u.GetHistoriesByUser(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("User.GetHistoriesByUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUser_CreateHeartbeats(t *testing.T) {
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
			if err := tt.u.CreateHeartbeats(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("User.CreateHeartbeats() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUser_GetDeviceStatus(t *testing.T) {
	type args struct {
		ctx context.Context
		d   *device.Device
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
			if err := tt.u.GetDeviceStatus(tt.args.ctx, tt.args.d); (err != nil) != tt.wantErr {
				t.Errorf("User.GetDeviceStatus() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUser_SetServiceParamsByUser(t *testing.T) {
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
			if err := tt.u.SetServiceParamsByUser(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("User.SetServiceParamsByUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUser_GetServiceParamsByUser(t *testing.T) {
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
			if err := tt.u.GetServiceParamsByUser(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("User.GetServiceParamsByUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUser_CheckConfigs(t *testing.T) {
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
			if err := tt.u.CheckConfigs(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("User.CheckConfigs() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
