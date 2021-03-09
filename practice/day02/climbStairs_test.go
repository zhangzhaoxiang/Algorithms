package day02

import "testing"

/*
@Time    : 2021/3/9 10:09 上午
@Author  : zx
@Email   : zhangzhaoxiang7@163.com
@File    : climbStairs_test.go.go
@Software: GoLand
*/

func Test_climbStairs(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "normal",
			args: args{n: 2},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairs(tt.args.n); got != tt.want {
				t.Errorf("climbStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
