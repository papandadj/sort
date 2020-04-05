package sort

import "testing"

func TestStalin(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name         string
		args         args
		wantIsOrders bool
	}{
		{
			"两个参数",
			args{
				arr: []int{20, 10},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotIsOrders := Stalin(tt.args.arr...); gotIsOrders != tt.wantIsOrders {
				t.Errorf("Stalin() = %v, want %v", gotIsOrders, tt.wantIsOrders)
			}
		})
	}
}
