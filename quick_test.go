package sort

import (
	"reflect"
	"testing"
)

func TestQuick(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"一个参数",
			args{
				arr: []int{20},
			},
			[]int{20},
		},
		{
			"两个参数",
			args{
				arr: []int{20, 10},
			},
			[]int{10, 20},
		}, {
			"四个参数",
			args{
				arr: []int{20, 10, 90, 40},
			},
			[]int{10, 20, 40, 90},
		}, {
			"6个参数",
			args{
				arr: []int{4, 30, 7, 9, 50, 80},
			},
			[]int{4, 7, 9, 30, 50, 80},
		}, {
			"10个参数",
			args{
				arr: []int{10, 1, 35, 61, 89, 36, 55, 300, 100, 200},
			},
			[]int{1, 10, 35, 36, 55, 61, 89, 100, 200, 300},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Quick(tt.args.arr...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Quick() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_quick(t *testing.T) {
	type args struct {
		arr   []int
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			quick(tt.args.arr, tt.args.left, tt.args.right)
		})
	}
}
