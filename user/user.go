package user

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-benchmark/scenario/device"

	http "github.com/gobench-io/gobench/clients/http"

	"github.com/bxcodec/faker/v3"
)

// Accesstoken access token
type Accesstoken struct {
	ID     string
	UserID string
}

// UserAccount user account
type UserAccount struct {
	Email    string `json:"email" faker:"email"`
	Password string `json:"password" faker:"password"`
}

// Home home
type Home struct {
	ID            string `json:"id,omitempty" faker:"-"`
	Name          string `json:"name" faker:"username"`
	CustomerEmail string `json:"customerEmail" faker:"email"`
}

// SignUp user sign up
func (u *User) SignUp(ctx context.Context) (err error) {
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	signupPath := "/api/users"
	signupURL := fmt.Sprintf("https://%s%s", u.host, signupPath)

	signupClient, err := http.NewHttpClient(ctx, signupPath)

	uaS, _ := json.Marshal(u.ua)

	// create new user
	if _, err = signupClient.Post(ctx, signupURL, uaS, headers); err != nil {
		return
	}

	return
}

// Login user log in
func (u *User) Login(ctx context.Context) (err error) {

	var at Accesstoken

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	loginPath := "/api/users/login"
	loginURL := fmt.Sprintf("https://%s%s", u.host, loginPath)
	loginClient, err := http.NewHttpClient(ctx, loginPath)

	uaS, _ := json.Marshal(u.ua)

	// login
	buf, err := loginClient.Post(ctx, loginURL, uaS, headers)
	if err != nil {
		return
	}

	if err = json.Unmarshal(buf, &at); err != nil {
		return
	}

	u.at = at

	return
}

// GetUserAccount get user account
func (u *User) GetUserAccount(ctx context.Context) (err error) {
	ua := UserAccount{}
	if err = faker.FakeData(&ua); err != nil {
		return
	}

	u.ua = ua
	return
}

func (u *User) headers() map[string]string {
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	if u.at.ID != "" {
		headers["Authorization"] = u.at.ID
	}
	return headers
}

// CreateSsEngine create softsecurity engine
func (u *User) CreateSsEngine(ctx context.Context) error {
	return u.createEngine(ctx, softSecurity)
}

// CreateLoEngine create location engine
func (u *User) CreateLoEngine(ctx context.Context) error {
	return u.createEngine(ctx, location)
}

// createEngine create an et engine for the user
// save the engine to local engines array
func (u *User) createEngine(ctx context.Context, et engineType) (err error) {
	newEnginePath := "/api/users/[id]/services"
	newEngineURL := fmt.Sprintf("https://%s/api/users/%s/services", u.host, u.at.UserID)
	newEngineClient, err := http.NewHttpClient(ctx, newEnginePath)

	req, _ := json.Marshal(map[string]engineType{
		"engineType": et,
	})

	var buf []byte
	if buf, err = newEngineClient.Post(ctx, newEngineURL, req, u.headers()); err != nil {
		return
	}

	var en engine
	if err = json.Unmarshal(buf, &en); err != nil {
		return
	}
	u.engines[et] = en

	return
}

// AttachDevice a user attach a device to a user
func (u *User) AttachDevice(ctx context.Context, d *device.Device) (err error) {
	// create new home
	// newDevicePath := "/api/users/[id]/devices"
	// newDeviceURL := fmt.Sprintf("https://%s/api/users/%s/devices", u.host, u.at.UserID)
	// newDeviceClient, err := http.NewHttpClient(ctx, newDevicePath)

	// req, _ := json.Marshal(map[string]string{
	// 	"deviceId":     d.ID,
	// 	"macAddress":   "a random mac address",
	// 	"wifiSsid":     "ssid of wifi",
	// 	"wifiPassword": "password of wifi",
	// })

	// if buf, err = newDeviceClient.Post(ctx, newDeviceURL, req, u.headers()); err != nil {
	// 	return
	// }

	return
}

// CreateDeviceSet 2.3 Create a deviceset
func (u *User) CreateDeviceSet(ctx context.Context) (err error) {
	// create new deviceset
	newDeviceSetPath := "/api/users/[id]/devicesets"
	newDeviceSetURL := fmt.Sprintf("https://%s/api/users/%s/devicesets", u.host, u.at.UserID)
	newDeviceSetClient, err := http.NewHttpClient(ctx, newDeviceSetPath)

	var name string
	_ = faker.FakeData(&name)
	req, _ := json.Marshal(map[string]string{
		"name": name,
	})

	var buf []byte
	if buf, err = newDeviceSetClient.Post(ctx, newDeviceSetURL, req, u.headers()); err != nil {
		return
	}
	var ds DeviceSet
	if err = json.Unmarshal(buf, &ds); err != nil {
		return
	}
	// save deviceset to user
	u.deviceSets[ds.ID] = ds
	return
}

// AddDevicesToDeviceSets 2.4 Add devices to devicesets, add 1 MO, 2 bots to the deviceset
func (u *User) AddDevicesToDeviceSets(ctx context.Context, d *device.Device) (err error) {
	if d == nil {
		err = fmt.Errorf("device is nil")
		return
	}
	for k, ds := range u.deviceSets {
		if err = u.AddDevicesToDeviceSet(ctx, &ds, *d); err != nil {
			return
		}
		// update deviceset
		u.deviceSets[k] = ds
	}
	return
}

// AddServicesToDeviceSets 2.4 Add devices to devicesets, add 1 MO, 2 bots to the deviceset
func (u *User) AddServicesToDeviceSets(ctx context.Context) (err error) {
	for _, ds := range u.deviceSets {
		if err = u.AddServiceToDeviceSets(ctx, ds); err != nil {
			return
		}
	}
	return
}

// GetHistoriesByUser get histories data of all services for a user
func (u *User) GetHistoriesByUser(ctx context.Context) (err error) {
	for _, v := range u.services {
		s := v

		if err = u.GetHistories(ctx, s); err != nil {
			return
		}
	}
	return
}

// CreateHeartbeats get real time data of all services for a user
func (u *User) CreateHeartbeats(ctx context.Context) (err error) {
	for _, v := range u.services {
		s := v

		if err = u.CreateHeartbeat(ctx, s); err != nil {
			return
		}
	}
	return
}

// GetDeviceStatus get status of a devive by a user
func (u *User) GetDeviceStatus(ctx context.Context, d *device.Device) (err error) {

	newDevicePath := "/api/devices/[id]/status"
	newDeviceURL := fmt.Sprintf("https://%s/api/devices/%s/status", u.host, d.ID)
	newDeviceClient, err := http.NewHttpClient(ctx, newDevicePath)

	if _, err = newDeviceClient.Get(ctx, newDeviceURL, u.headers()); err != nil {
		return
	}
	return
}

// SetServiceParamsByUser get service params of service by a user
func (u *User) SetServiceParamsByUser(ctx context.Context) (err error) {
	for _, v := range u.services {
		s := v

		if err = u.SetServiceParams(ctx, s); err != nil {
			return
		}
	}
	return
}

// GetServiceParamsByUser get service params of service by a user
func (u *User) GetServiceParamsByUser(ctx context.Context) (err error) {
	for _, v := range u.services {
		s := v

		if err = u.GetServiceParams(ctx, s); err != nil {
			return
		}
	}
	return
}

func (u *User) CheckConfigs(ctx context.Context) (err error) {
	for _, v := range u.services {
		s := v

		if err = u.CheckConfig(ctx, s); err != nil {
			return
		}
	}
	return
}
