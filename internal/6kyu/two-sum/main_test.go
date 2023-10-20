package twosum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	type args struct {
		numbers []int
		target  int
	}
	tests := []struct {
		name string
		args args
		want [2]int
	}{
		{"", args{[]int{1, 2, 3}, 4}, [2]int{0, 2}},
		{"", args{[]int{1234, 5678, 9012}, 14690}, [2]int{1, 2}},
		{"", args{[]int{2, 2, 3}, 4}, [2]int{0, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TwoSum(tt.args.numbers, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TwoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
