package problems

import (
	"reflect"
	"testing"
)

/*
@Time    : 2021/3/8 8:44 下午
@Author  : zx
@Email   : zhangzhaoxiang7@163.com
@File    : twoSum_test.go.go
@Software: GoLand
*/

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "normal",
			args: args{
				nums:   []int{1, 2, 3, 4, 5, 6, 7, 8},
				target: 15,
			},
			want: []int{6, 7},
		},
		{
			name: "nil",
			args: args{
				nums:   nil,
				target: 10,
			},
			want: nil,
		},
		{
			name: "empty",
			args: args{
				nums:   []int{},
				target: 10,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
