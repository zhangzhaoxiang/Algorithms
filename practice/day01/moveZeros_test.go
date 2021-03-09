package problems

import (
	"reflect"
	"testing"
)

/*
@Time    : 2021/3/8 7:00 下午
@Author  : zx
@Email   : zhangzhaoxiang7@163.com
@File    : moveZeros_test.go.go
@Software: GoLand
*/

func Test_moveZeros(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "normal",
			args: struct{ nums []int }{nums: []int{1, 0, 2, 0, 3}},
			want: []int{1, 2, 3, 0, 0},
		},
		{
			name: "nil",
			args: struct{ nums []int }{nums: nil},
			want: nil,
		},
		{
			name: "empty",
			args: struct{ nums []int }{nums: []int{}},
			want: []int{},
		},
		{
			name: "zeros",
			args: struct{ nums []int }{nums: []int{0, 0, 0, 0, 0}},
			want: []int{0, 0, 0, 0, 0},
		},
		{
			name: "noZeros",
			args: struct{ nums []int }{nums: []int{1, 2, 3, 4, 5}},
			want: []int{1, 2, 3, 4, 5},
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := moveZeros(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("moveZeros() = %v, want %v", got, tt.want)
			}
		})
	}
}
