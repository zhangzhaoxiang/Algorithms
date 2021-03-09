package problems

import "testing"

/*
@Time    : 2021/3/8 8:20 下午
@Author  : zx
@Email   : zhangzhaoxiang7@163.com
@File    : maxArea_test.go.go
@Software: GoLand
*/

func Test_maxArea(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "normal",
			args: args{height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7}},
			want: 49,
		},
		{
			name: "nil",
			args: args{height: nil},
			want: 0,
		},
		{
			name: "empty",
			args: args{height: []int{}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxArea(tt.args.height); got != tt.want {
				t.Errorf("maxArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
