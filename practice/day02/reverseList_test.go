package day02

import (
	. "Algorithms/pkg"
	"reflect"
	"testing"
)

/*
@Time    : 2021/3/9 4:54 下午
@Author  : zx
@Email   : zhangzhaoxiang7@163.com
@File    : reverseList_test.go.go
@Software: GoLand
*/

func Test_reverseList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "normal",
			args: args{head: ListConvertListNode([]int{1, 2, 3, 4, 5})},
			want: ListConvertListNode([]int{5, 4, 3, 2, 1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}
