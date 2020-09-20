package user

import (
	"reflect"
	"testing"

	"github.com/go-benchmark/scenario/config"
)

func TestNewUser(t *testing.T) {
	type args struct {
		opts *config.Options
	}
	tests := []struct {
		name string
		args args
		want *User
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUser(tt.args.opts); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
