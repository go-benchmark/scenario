package user

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/gobench-io/gobench/clients/http"
)

// Error
var (
	ErrEngineNotSupport = errors.New("engine type not supported")
	ErrNoBot            = errors.New("get histories for location service require at least bot")
)

// FSService represents to fs service
type FSService struct {
	ID          string `json:"id,omitempty" faker:"-"`
	Name        string `json:"name,omitempty" faker:"name"`
	EngineType  string `json:"engineType"`
	UserID      string `json:"userId"`
	DeviceSetID string `json:"devicesetId,omitempty"`
}

// ConfigService represent to get config service response
type ConfigService struct {
	Origin string        `json:"origin"`
	Config ServiceStatus `json:"config"`
}

// ServiceStatus service run status
type ServiceStatus struct {
	Timestamp  int    `json:"timestamp,omitempty"`
	EngineType string `json:"engineType,omitempty"`
	ServiceID  string `json:"serviceId,omitempty"`
	CMD        string `json:"cmd,omitempty"`
	Success    bool   `json:"success,omitempty"`
}

// ServiceParams represent to get/set params
type ServiceParams struct {
	EngineType string `json:"engineType"`
	ServiceID  string `json:"serviceId"`
	EM         EM     `json:"engineMsg"`
}

// LocationEM em for location
type LocationEM struct {
	Saving         bool `json:"saving"`
	RealtimeLength int  `json:"realtimeLength"`
}

//EM engine message represent to optional params of a service
type EM struct {
	Saving               bool            `json:"saving"`
	RealtimeLength       int             `faker:"oneof: 6, 30" json:"realtimeLength"`
	NotificationInterval int             `faker:"oneof: 10, 30" json:"notificationInterval,omitempty"`
	OverallSensitivity   int             `faker:"boundary_start=0, boundary_end=100" json:"overallSensitivity,omitempty"`
	Schedules            []schedule      `json:"schedules,omitempty"`
	ManualSchedule       *manualSchedule `json:"manualSchedule,omitempty"`
}
type schedule struct {
	TimeInstant []int `faker:"boundary_start=1595414207, boundary_end=1595474207" faker:"slice_len=4" json:"timeInstant"`
	Active      bool  `json:"active"`
}
type manualSchedule struct {
	Alert bool  `json:"alert"`
	At    int64 `faker:"unix_time" json:"at"`
}

// CreateService create a service by engine type
func (u *User) CreateService(ctx context.Context, ds DeviceSet, et engineType) (err error) {
	newServicePath := "/api/devicesets/[id]/services"
	newServiceURL := fmt.Sprintf("https://%s/api/devicesets/%s/services", u.host, ds.ID)
	newServiceClient, err := http.NewHttpClient(ctx, newServicePath)

	req, _ := json.Marshal(map[string]string{
		"name":       faker.Name(),
		"engineType": string(et),
	})
	var buf []byte
	if buf, err = newServiceClient.Post(ctx, newServiceURL, req, u.headers()); err != nil {
		return
	}
	var s FSService
	if err = json.Unmarshal(buf, &s); err != nil {
		return
	}
	u.services[s.ID] = s

	return
}

// AttachServices attach services to a deviceset
func (u *User) AttachServices(ctx context.Context, ds DeviceSet, serviceIds []string) (err error) {
	newServicePath := "/api/devicesets/[id]/attach-services"
	newServiceURL := fmt.Sprintf("https://%s/api/devicesets/%s/attach-services", u.host, ds.ID)
	newServiceClient, err := http.NewHttpClient(ctx, newServicePath)

	req, _ := json.Marshal(map[string][]string{
		"serviceIds": serviceIds,
	})
	var buf []byte
	if buf, err = newServiceClient.Post(ctx, newServiceURL, req, u.headers()); err != nil {
		return
	}
	var s FSService
	if err = json.Unmarshal(buf, &s); err != nil {
		return
	}
	u.services[s.ID] = s

	return
}

// StartServices start a service then get config to check
func (u *User) StartServices(ctx context.Context) (err error) {
	return u.RunServices(ctx, CmdStart)
}

// StopServices start a service then get config to check
func (u *User) StopServices(ctx context.Context) (err error) {
	return u.RunServices(ctx, CmdStop)
}

// RunServices run services
func (u *User) RunServices(ctx context.Context, cmd string) (err error) {
	var wg sync.WaitGroup
	c := make(chan error)
	// start services
	for _, v := range u.services {
		wg.Add(1)
		s := v

		go func(s FSService) {
			defer wg.Done()
			defer func() {
				c <- err
			}()
			// get deviceset from service
			ds, ok := u.deviceSets[s.DeviceSetID]
			if !ok {
				err = fmt.Errorf("service %s dont have deviceset", s.ID)
				return
			}
			// get devices from deviceset
			bots := make([]string, 0, len(ds.devices))
			var moID string
			for _, v := range ds.devices {
				d := v

				if d.Type == DeviceMo {
					if moID != "" {
						continue
					}

					moID = d.ID
					continue
				}
				bots = append(bots, d.ID)
			}
			// require at least mo device from deviceset
			if moID == "" {
				err = fmt.Errorf("service %s require at least mo device", s.ID)
				return
			}
			// start service
			err = u.RunService(ctx, s, cmd, moID, bots)

			return
		}(s)
	}
	// check run service
	for range u.services {
		select {
		case err := <-c:
			if err != nil {
				return err
			}
		}
	}
	wg.Wait()
	return
}

// RunService run services
func (u *User) RunService(ctx context.Context, s FSService, cmd, moID string, bots []string) (err error) {

	newServicePath := "/api/services/run/[deviceId(moId)]"
	newServiceURL := fmt.Sprintf("https://%s/api/services/run/%s", u.host, moID)
	newServiceClient, err := http.NewHttpClient(ctx, newServicePath)

	var em EM
	if err = faker.FakeData(&em); err != nil {
		return
	}
	em.RealtimeLength = u.heartbeatInterval

	req, _ := json.Marshal(map[string]interface{}{
		"engineType":    s.EngineType,
		"serviceId":     s.ID,
		"cmd":           cmd,
		"bots":          bots,
		"operationInfo": em,
	})

	if _, err = newServiceClient.Post(ctx, newServiceURL, req, u.headers()); err != nil {
		return
	}
	return
}

// CheckConfig represent to check config after run service
func (u *User) CheckConfig(ctx context.Context, s FSService) (err error) {
	var cs ConfigService

	cs, err = u.GetConfig(ctx, s)
	if err != nil {
		return
	}
	if !cs.Config.Success {
		err = fmt.Errorf("service %s run failed", s.ID)
		return
	}
	return
}

// GetConfig run services
func (u *User) GetConfig(ctx context.Context, s FSService) (cs ConfigService, err error) {

	newServicePath := "/api/services/[id]/config"
	newServiceURL := fmt.Sprintf("https://%s/api/services/%s/config", u.host, s.ID)
	newServiceClient, err := http.NewHttpClient(ctx, newServicePath)

	var buf []byte
	if buf, err = newServiceClient.Get(ctx, newServiceURL, u.headers()); err != nil {
		return
	}
	if err = json.Unmarshal(buf, &cs); err != nil {
		return
	}
	return
}

// GetHistories get histories of a service
func (u *User) GetHistories(ctx context.Context, s FSService) (err error) {

	newServicePath := "/api/services/[id]/histories"
	newServiceURL := fmt.Sprintf("https://%s/api/services/%s/histories", u.host, s.ID)
	newServiceClient, err := http.NewHttpClient(ctx, newServicePath)

	switch s.EngineType {
	case string(location):
		// get deviceset from service
		ds, ok := u.deviceSets[s.DeviceSetID]
		if !ok {
			err = fmt.Errorf("service %s dont have deviceset", s.ID)
			return
		}
		// random get histories for a zone or all devices.
		var bots string
		rand.Seed(time.Now().UnixNano())
		randHistory := rand.Intn(len(u.services))
		randZone := rand.Intn(len(ds.zones))
		switch randHistory {
		// get histories for a zone
		case 1:
			i := 0
			for _, v := range ds.zones {
				z := v
				if i != randZone {
					i++
					continue
				}
				for _, deviceID := range z.DeviceIDs {
					if bots == "" {
						bots = deviceID
						continue
					}
					bots += "," + deviceID
				}
				break
			}
			break
			// get histories for all zones
		default:
			for _, v := range ds.devices {
				d := v

				if bots == "" {
					bots = d.ID
					continue
				}
				bots += "," + d.ID
			}
			break
		}
		// check bots
		if len(bots) == 0 {
			err = ErrNoBot
			return
		}
		// begin := time.Now().AddDate(0, 0, -1).Unix()
		// end := time.Now().Unix()
		newServiceURL += fmt.Sprintf("?devices=%s", bots)
		break
	default:
		break
	}
	if _, err = newServiceClient.Get(ctx, newServiceURL, u.headers()); err != nil {
		return
	}
	return
}

// CreateHeartbeat get realtime data of a service
func (u *User) CreateHeartbeat(ctx context.Context, s FSService) (err error) {

	newServicePath := "/api/services/[id]/heartbeats"
	newServiceURL := fmt.Sprintf("https://%s/api/services/%s/heartbeats", u.host, s.ID)
	newServiceClient, err := http.NewHttpClient(ctx, newServicePath)

	if _, err = newServiceClient.Post(ctx, newServiceURL, nil, u.headers()); err != nil {
		return
	}
	return
}

// SetServiceParams set params of a service
// deprecated use set service params when start a service instead
func (u *User) SetServiceParams(ctx context.Context, s FSService) (err error) {

	newServicePath := "/api/services/[id]/params"
	newServiceURL := fmt.Sprintf("https://%s/api/services/%s/params", u.host, s.ID)
	newServiceClient, err := http.NewHttpClient(ctx, newServicePath)

	var em EM
	if err = faker.FakeData(&em); err != nil {
		return
	}
	em.RealtimeLength = u.heartbeatInterval
	sp := ServiceParams{
		EngineType: s.EngineType,
		ServiceID:  s.ID,
		EM:         em,
	}

	req, _ := json.Marshal(sp)
	if _, err = newServiceClient.Post(ctx, newServiceURL, req, u.headers()); err != nil {
		return
	}
	return
}

// GetServiceParams get params of a service
func (u *User) GetServiceParams(ctx context.Context, s FSService) (err error) {

	newServicePath := "/api/services/[id]/params"
	newServiceURL := fmt.Sprintf("https://%s/api/services/%s/params", u.host, s.ID)
	newServiceClient, err := http.NewHttpClient(ctx, newServicePath)

	if _, err = newServiceClient.Get(ctx, newServiceURL, u.headers()); err != nil {
		return
	}
	return
}
