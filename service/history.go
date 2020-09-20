package service

import (
	"encoding/json"
	"errors"
	"math/rand"
	"time"

	"github.com/bxcodec/faker/v3"
)

// Error
var (
	ErrEngineNotSupport              = errors.New("engine type not supported")
	ErrWellbeingHistoricalNotSupport = errors.New("engine wellbeing historical type not supported")
)

type baseH struct {
	Timestamp  int64      `json:"timestamp"`
	EngineType EngineType `json:"engineType"`
	ServiceID  string     `json:"serviceId"`
}

type database struct {
	Detection int `faker:"boundary_start=0, boundary_end=2" json:"detection"`
}
type wellbeingDecisionDatabase struct {
	DatabaseType    string        `json:"databaseType"`
	TimeToBed       int64         `json:"timeToBed,omitempty"`
	TimeBeforeSleep int           `json:"timeBeforeSleep,omitempty"`
	WakeUpCount     int           `json:"wakeUpCount,omitempty"`
	TimeInBed       int           `json:"timeInBed,omitempty"`
	TimeInRem       int           `json:"timeInRem,omitempty"`
	TimeInNrem      int           `json:"timeInNrem,omitempty"`
	SleepScore      int           `json:"sleepScore,omitempty"`
	Stages          []int         `json:"stages,omitempty"`
	TimeInStages    []int         `json:"timeInStages,omitempty"`
	MotionStates    []motionState `json:"motionStates,omitempty"`
}
type wellbeingDeletionDatabase struct {
	Key string `json:"key"`
}
type motionState struct {
	Tme       int    `json:"tme"`
	Detection int    `json:"detection"`
	Bot       string `json:"bot"`
}
type wellbeingHistorical struct {
	Type func(wH, WellbeingHistoricalType) (wH, error)
}
type locationDatabase struct {
	Locations []location `json:"locations"`
        Detection int `json:"detection"` // backward compatibility
}
type location struct {
	Detection int    `faker:"boundary_start=0, boundary_end=2" json:"detection"`
	DeviceID  string `json:"deviceId"` //bot id
}
type ssH struct {
	baseH
	Database database `json:"database"`
}
type lH struct {
	baseH
	Database locationDatabase `json:"database"`
}
type wH struct {
	baseH
	Database *wellbeingDecisionDatabase `json:"database,omitempty"`
	Delete   *wellbeingDeletionDatabase `json:"delete,omitempty"`
}
type historyFn func(string, EngineType) ([]byte, error)

// History handle history response by engine
func (s *Service) History() (result []byte, err error) {
	switch s.EngineType {
	case SoftSecurity:
		return s.ssHistoryv2()
	case Wellbeing:
		return s.wellbeingHistory()
	case Location:
		return s.locationHistory()
	default:
		err = ErrEngineNotSupport
		return
	}
}

// create a random softSecurity history
func (s *Service) ssHistory() (result []byte, err error) {
	h := ssH{
		baseH: baseH{
			Timestamp:  time.Now().Unix(),
			EngineType: s.EngineType,
			ServiceID:  s.ServiceID,
		},
	}
	if err = faker.FakeData(&h.Database); err != nil {
		return
	}

	result, err = json.Marshal(h)

	return
}
func (s *Service) locationHistory() (result []byte, err error) {

	h := lH{
		baseH: baseH{
			Timestamp:  time.Now().Unix(),
			EngineType: s.EngineType,
			ServiceID:  s.ServiceID,
		},
		Database: locationDatabase{
			Locations: []location{},
		},
	}
	for _, v := range s.Bots {

		l := location{
			Detection: rand.Intn(2),
			DeviceID:  v,
		}
		h.Database.Locations = append(h.Database.Locations, l)
	}
	result, err = json.Marshal(h)
	return
}
func (s *Service) wellbeingHistory() (result []byte, err error) {

	h := wH{
		baseH: baseH{
			Timestamp:  time.Now().Unix(),
			EngineType: s.EngineType,
			ServiceID:  s.ServiceID,
		},
	}
	// random wellbeing historical output
	wTypes := []WellbeingHistoricalType{
		WellbeingSleep,
		WellbeingLocalization,
		WellbeingDeletion,
	}
	wType := wTypes[rand.Intn(len(wTypes))]

	h, err = wellbeingPayload(h, wType)
	if err != nil {
		return
	}

	result, err = json.Marshal(h)
	return
}
func wellbeingPayload(h wH, wht WellbeingHistoricalType) (wH, error) {

	switch wht {
	case WellbeingSleep:
		if err := faker.FakeData(h.Database); err != nil {
			return h, err
		}
		h.Database.DatabaseType = string(wht)
		h.Database.MotionStates = nil
		return h, nil
	case WellbeingLocalization:
		if err := faker.FakeData(h.Database); err != nil {
			return h, err
		}
		ms := h.Database.MotionStates
		h.Database = &wellbeingDecisionDatabase{
			MotionStates: ms,
			DatabaseType: string(wht),
		}
		return h, nil
	case WellbeingDeletion:
		if err := faker.FakeData(h.Delete); err != nil {
			return h, err
		}
		return h, nil
	default:
		return h, ErrWellbeingHistoricalNotSupport
	}
}
func (s *Service) ssHistoryv2() (result []byte, err error) {

	h := lH{
		baseH: baseH{
			Timestamp:  time.Now().Unix(),
			EngineType: s.EngineType,
			ServiceID:  s.ServiceID,
		},
		Database: locationDatabase{
			Locations: []location{},
		},
	}
	for _, v := range s.Bots {

		l := location{
			Detection: rand.Intn(2),
			DeviceID:  v,
		}
		h.Database.Locations = append(h.Database.Locations, l)
	}
	result, err = json.Marshal(h)
	return
}
