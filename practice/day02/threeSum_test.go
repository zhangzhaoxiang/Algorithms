package day02

import (
	"reflect"
	"testing"
)

/*
@Time    : 2021/3/9 3:30 下午
@Author  : zx
@Email   : zhangzhaoxiang7@163.com
@File    : threeSum_test.go.go
@Software: GoLand
*/

func Test_threeSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "normal",
			args: args{nums: []int{-1, 0, 1, 2, -1, -4}},
			want: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name: "empty",
			args: args{nums: []int{}},
			want: nil,
		},
		{
			name: "one",
			args: args{nums: []int{0}},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSum(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("threeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
