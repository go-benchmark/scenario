package service

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHistorySoftSecurity(t *testing.T) {
	s := Service{
		EngineType: SoftSecurity,
		ServiceID:  "0123",
	}

	r, err := s.History()
	assert.Nil(t, err)

	h := ssH{}

	assert.Nil(t, json.Unmarshal(r, &h))
	assert.Equal(t, SoftSecurity, h.EngineType)
	assert.Equal(t, "0123", h.ServiceID)
	assert.Contains(t, []int{0, 1, 2}, h.Database.Detection)
}

func TestService_History(t *testing.T) {
	tests := []struct {
		name       string
		s          *Service
		wantResult []byte
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := tt.s.History()
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.History() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("Service.History() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestService_ssHistory(t *testing.T) {
	tests := []struct {
		name       string
		s          *Service
		wantResult []byte
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := tt.s.ssHistory()
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.ssHistory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("Service.ssHistory() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestService_locationHistory(t *testing.T) {
	tests := []struct {
		name       string
		s          *Service
		wantResult []byte
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := tt.s.locationHistory()
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.locationHistory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("Service.locationHistory() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestService_wellbeingHistory(t *testing.T) {
	tests := []struct {
		name       string
		s          *Service
		wantResult []byte
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := tt.s.wellbeingHistory()
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.wellbeingHistory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("Service.wellbeingHistory() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func Test_wellbeingPayload(t *testing.T) {
	type args struct {
		h   wH
		wht WellbeingHistoricalType
	}
	tests := []struct {
		name    string
		args    args
		want    wH
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := wellbeingPayload(tt.args.h, tt.args.wht)
			if (err != nil) != tt.wantErr {
				t.Errorf("wellbeingPayload() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("wellbeingPayload() = %v, want %v", got, tt.want)
			}
		})
	}
}
