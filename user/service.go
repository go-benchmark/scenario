package user

import (
	"context"

	"github.com/go-benchmark/scenario/config"
	"github.com/go-benchmark/scenario/device"
)

// User represents to user model
type User struct {
	host              string
	ua                UserAccount
	at                Accesstoken
	home              Home
	heartbeatInterval int

	engines    map[engineType]engine
	deviceSets map[string]DeviceSet
	services   map[string]FSService
}

// UserServicer represents to user interface
type UserServicer interface {
	GetUserAccount(ctx context.Context) (err error)
	SignUp(ctx context.Context) (err error)
	Login(ctx context.Context) (err error)
	CreateSsEngine(ctx context.Context) error
	CreateLoEngine(ctx context.Context) error
	AttachDevice(ctx context.Context, d device.Device) (err error)
	CreateDeviceSet(ctx context.Context) (err error)
	AddDevicesToDeviceSets(ctx context.Context, d *device.Device) (err error)
	AddServicesToDeviceSets(ctx context.Context) (err error)
	AddZones(ctx context.Context, zoneCount int) (err error)
	RunServices(ctx context.Context) (err error)
	StartServices(ctx context.Context) (err error)
	StopServices(ctx context.Context) (err error)
	GetHistories(context.Context, FSService) (err error)
	GetHistoriesByUser(ctx context.Context) (err error)
	CreateHeartbeat(context.Context, FSService) (err error)
	CreateHeartbeats(ctx context.Context) (err error)
	GetDeviceStatus(context.Context, *device.Device) (err error)
	GetServiceParamsByUser(ctx context.Context) (err error)

	// mixins
	createEngine(ctx context.Context, et engineType) (err error)
	AddDevicesToDeviceSet(ctx context.Context, ds *DeviceSet, d device.Device) (err error)
	AddServiceToDeviceSets(ctx context.Context, ds DeviceSet) (err error)
	CreateService(ctx context.Context, ds DeviceSet, et engineType) (err error)
	AttachServices(ctx context.Context, ds DeviceSet) (err error)
	AddZone(ctx context.Context, ds DeviceSet) (err error)
	RunService(ctx context.Context, s FSService, cmd, moID string, bots []string) (err error)
	GetConfig(ctx context.Context, s FSService) (err error)
	GetServiceParams(ctx context.Context, s FSService) (err error)
}

// NewUser init user service with default
func NewUser(opts *config.Options) *User {
	return &User{
		host:              opts.Host,
		heartbeatInterval: opts.UC.RealtimeHBInterval,
		engines:           make(map[engineType]engine),
		deviceSets:        make(map[string]DeviceSet),
		services:          make(map[string]FSService),
	}
}
