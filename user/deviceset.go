package user

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/bxcodec/faker/v3"
	"github.com/go-benchmark/scenario/device"
	"github.com/gobench-io/gobench/clients/http"
)

// DeviceSet represent to deviceset
type DeviceSet struct {
	ID   string `json:"id"`
	Name string `json:"name"`

	devices map[string]device.DeviceReq
	zones   map[string]Zone
}

// Zone represent to zone deviceset
type Zone struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	UserID      string   `json:"userId"`
	DevicesetID string   `json:"devicesetId"`
	DeviceIDs   []string `json:"deviceIds"`
}

// NewDeviceSet create a device mqtt credential
func (u *User) NewDeviceSet(ctx context.Context) (ds DeviceSet, err error) {

	ds.devices = make(map[string]device.DeviceReq)
	return
}

// AddDevicesToDeviceSet 2.4 Add devices to a deviceset, add 1 MO, 2 bots to the deviceset
func (u *User) AddDevicesToDeviceSet(ctx context.Context, ds *DeviceSet, d device.Device) (err error) {
	// add devices to a deviceset
	newDeviceSetPath := "/api/devicesets/[id]/add-devices"
	newDeviceSetURL := fmt.Sprintf("https://%s/api/devicesets/%s/add-devices", u.host, ds.ID)
	newDeviceSetClient, err := http.NewHttpClient(ctx, newDeviceSetPath)

	req, _ := json.Marshal([]device.DeviceReq{
		// 1 mo device
		{
			ID:           d.ID,
			Type:         DeviceMo,
			Name:         faker.Name(),
			MacAddress:   faker.MacAddress(),
			WifiSsid:     faker.UUIDHyphenated(),
			WifiPassword: faker.Password(),
		},
		// 2 bots
		{
			ID:           faker.UUIDHyphenated(),
			Type:         DeviceBot,
			Name:         faker.Name(),
			MacAddress:   faker.MacAddress(),
			WifiSsid:     faker.UUIDHyphenated(),
			WifiPassword: faker.Password(),
		},
		{
			ID:           faker.UUIDHyphenated(),
			Type:         DeviceBot,
			Name:         faker.Name(),
			MacAddress:   faker.MacAddress(),
			WifiSsid:     faker.UUIDHyphenated(),
			WifiPassword: faker.Password(),
		},
	})

	var buf []byte
	if buf, err = newDeviceSetClient.Post(ctx, newDeviceSetURL, req, u.headers()); err != nil {
		return
	}
	var devices []device.DeviceReq
	if err = json.Unmarshal(buf, &devices); err != nil {
		return
	}
	if ds.devices == nil {
		ds.devices = make(map[string]device.DeviceReq)
	}
	// save devices to deviceset
	for _, device := range devices {
		ds.devices[device.ID] = device
	}
	return
}

//AddServiceToDeviceSets 2.5 Add services to a deviceset
func (u *User) AddServiceToDeviceSets(ctx context.Context, ds DeviceSet) (err error) {

	for _, ds := range u.DeviceSets {
		// create 2 service with engine type in [softSecurity,location]
		if err = u.CreateService(ctx, ds, softSecurity); err != nil {
			return
		}
	}
	return
}

//AddZones ### 2.6. Create 2 zones, a zone has one bot
func (u *User) AddZones(ctx context.Context, zoneCount int) (err error) {
	for k, ds := range u.DeviceSets {
		for i := 0; i < zoneCount; i++ {
			if err = u.AddZone(ctx, &ds); err != nil {
				return
			}
		}
		u.DeviceSets[k] = ds
	}
	return
}

// AddZone add zone and attach device(bot) into it
func (u *User) AddZone(ctx context.Context, ds *DeviceSet) (err error) {
	newDeviceSetPath := "/api/devicesets/[id]/zones"
	newDeviceSetURL := fmt.Sprintf("https://%s/api/devicesets/%s/zones", u.host, ds.ID)
	newDeviceSetClient, err := http.NewHttpClient(ctx, newDeviceSetPath)

	// get devices(bots) from deviceset
	deviceIds := make([]string, 0, 1) // in this case, a zone has 1 bot
	deviceOccupied := make(map[string]struct{})
	for _, v := range ds.zones {
		for _, deviceID := range v.DeviceIDs {
			deviceOccupied[deviceID] = struct{}{}
		}
	}

	for _, v := range ds.devices {
		d := v
		if d.Type == DeviceMo {
			continue
		}
		if _, ok := deviceOccupied[d.ID]; ok {
			// this device already on another zone
			continue
		}
		deviceIds = append(deviceIds, d.ID)
		break
	}

	req, _ := json.Marshal(map[string]interface{}{
		"name":      faker.Name(),
		"deviceIds": deviceIds,
	})
	var buf []byte
	if buf, err = newDeviceSetClient.Post(ctx, newDeviceSetURL, req, u.headers()); err != nil {
		return
	}
	var zone Zone
	if err = json.Unmarshal(buf, &zone); err != nil {
		return
	}
	zone.DeviceIDs = deviceIds
	if ds.zones == nil {
		ds.zones = make(map[string]Zone)
	}
	ds.zones[zone.ID] = zone

	return
}
