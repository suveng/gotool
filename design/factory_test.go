package design

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_build(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    Restaurant
		wantErr bool
	}{
		//  Add test cases.
		{
			name: "1",
			args: args{
				name: "1",
			},
		},
		{
			name: "2",
			args: args{
				name: "2",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := build(tt.args.name)
			fmt.Println(got)
			assert.True(t, err == nil, "build error")
			assert.True(t, got != nil, "not true")
		})
	}
}
