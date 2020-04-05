package sort

import (
	"reflect"
	"testing"
)

func Test_checkSorted(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				arr: []int{1, 2, 5, 6},
			},
			want: true,
		}, {
			args: args{
				arr: []int{1, 2, 5, 7, 8},
			},
			want: true,
		}, {
			args: args{
				arr: []int{1, 2, 5, 6, 3, 4},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkSorted(tt.args.arr...); got != tt.want {
				t.Errorf("checkSorted() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRandom(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
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
				arr: []int{4, 7, 9, 30, 50, 80},
			},
			[]int{4, 7, 9, 30, 50, 80},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Random(tt.args.arr...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Random() = %v, want %v", got, tt.want)
			}
		})
	}
}
