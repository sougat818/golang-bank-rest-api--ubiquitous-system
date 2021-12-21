package app

import (
	"reflect"
	"testing"
)

func Test_getHealthMessage(t *testing.T) {
	tests := []struct {
		name string
		want map[string]bool
	}{
		{"ok", map[string]bool{"ok": true}},
		{"nok", map[string]bool{"nok": true}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getHealthMessage(tt.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getHealthMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
