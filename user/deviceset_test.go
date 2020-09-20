package user

import (
	"context"
	"reflect"
	"testing"

	"github.com/go-benchmark/scenario/device"
)

func TestUser_NewDeviceSet(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		u       *User
		args    args
		wantDs  DeviceSet
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDs, err := tt.u.NewDeviceSet(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("User.NewDeviceSet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotDs, tt.wantDs) {
				t.Errorf("User.NewDeviceSet() = %v, want %v", gotDs, tt.wantDs)
			}
		})
	}
}

func TestUser_AddDevicesToDeviceSet(t *testing.T) {
	type args struct {
		ctx context.Context
		ds  *DeviceSet
		d   device.Device
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
			if err := tt.u.AddDevicesToDeviceSet(tt.args.ctx, tt.args.ds, tt.args.d); (err != nil) != tt.wantErr {
				t.Errorf("User.AddDevicesToDeviceSet() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUser_AddServiceToDeviceSets(t *testing.T) {
	type args struct {
		ctx context.Context
		ds  DeviceSet
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
			if err := tt.u.AddServiceToDeviceSets(tt.args.ctx, tt.args.ds); (err != nil) != tt.wantErr {
				t.Errorf("User.AddServiceToDeviceSets() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUser_AddZones(t *testing.T) {
	type args struct {
		ctx       context.Context
		zoneCount int
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
			if err := tt.u.AddZones(tt.args.ctx, tt.args.zoneCount); (err != nil) != tt.wantErr {
				t.Errorf("User.AddZones() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUser_AddZone(t *testing.T) {
	type args struct {
		ctx context.Context
		ds  *DeviceSet
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
			if err := tt.u.AddZone(tt.args.ctx, tt.args.ds); (err != nil) != tt.wantErr {
				t.Errorf("User.AddZone() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
