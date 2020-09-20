package main

import (
	"context"
	"reflect"
	"testing"

	"github.com/gobench-io/gobench/executor/scenario"
)

func TestExport(t *testing.T) {
	tests := []struct {
		name string
		want scenario.Vus
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := export(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Export() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_f(t *testing.T) {
	type args struct {
		ctx context.Context
		vui int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f(tt.args.ctx, tt.args.vui)
		})
	}
}
