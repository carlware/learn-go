package main

import "testing"

func Test_finalCostCalc(t *testing.T) {
	type args struct {
		closingStockCount int
		cost              int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := finalCostCalc(tt.args.closingStockCount, tt.args.cost); got != tt.want {
				t.Errorf("finalCostCalc() = %v, want %v", got, tt.want)
			}
		})
	}
}
