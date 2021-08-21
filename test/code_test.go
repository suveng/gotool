package test

import (
	"errors"
	"github.com/agiledragon/gomonkey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetPersonDetail(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name    string
		args    args
		want    *PersonDetail
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "invalid username", args: args{username: "steven xxx"}, want: nil, wantErr: true},
		{name: "invalid email", args: args{username: "invalid_email"}, want: nil, wantErr: true},
		{name: "throw err", args: args{username: "throw_err"}, want: nil, wantErr: true},
		{name: "valid return", args: args{username: "steven"}, want: &PersonDetail{Username: "steven", Email: "12345678@qq.com"}, wantErr: false},
	}

	// gomonkey打函数桩
	outputs := []gomonkey.OutputCell{
		{
			Values: gomonkey.Params{&PersonDetail{Username: "invalid email", Email: "test.com"}, nil},
		},
		{
			Values: gomonkey.Params{nil, errors.New("request error")},
		},
		{
			Values: gomonkey.Params{&PersonDetail{Username: "steven", Email: "12345678@qq.com"}, nil},
		},
	}
	// mock
	patches := gomonkey.ApplyFuncSeq(getPersonDetailRedis, outputs)
	// 执行完释放
	defer patches.Reset()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetPersonDetail(tt.args.username)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err != nil)

		})
	}
}
