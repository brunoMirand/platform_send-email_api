package campaign

import (
	"reflect"
	"testing"
)

func TestNewCampaign(t *testing.T) {
	type args struct {
		name    string
		content string
		emails  []string
	}
	tests := []struct {
		name string
		args args
		want *Campaign
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCampaign(tt.args.name, tt.args.content, tt.args.emails); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCampaign() = %v, want %v", got, tt.want)
			}
		})
	}
}
