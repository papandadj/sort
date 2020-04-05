package sort

import (
	"reflect"
	"testing"
)

func TestSleep(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name       string
		args       args
		wantOrders []int
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
				arr: []int{4, 30, 7, 9, 50, 80},
			},
			[]int{4, 7, 9, 30, 50, 80},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOrders := Sleep(tt.args.arr...); !reflect.DeepEqual(gotOrders, tt.wantOrders) {
				t.Errorf("Sleep() = %v, want %v", gotOrders, tt.wantOrders)
			}
		})
	}
}
