package directionsreduction

import (
	"reflect"
	"testing"
)

func TestDirReduc(t *testing.T) {
	type args struct {
		arr []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"1 empty result", args{arr: []string{"NORTH", "SOUTH", "EAST", "WEST"}}, []string{}},
		{"2 one", args{arr: []string{"NORTH", "SOUTH", "SOUTH", "EAST", "WEST", "NORTH", "WEST"}}, []string{"WEST"}},
		{"3 not reduce", args{arr: []string{"NORTH", "WEST", "SOUTH", "EAST"}}, []string{"NORTH", "WEST", "SOUTH", "EAST"}},
		{"4", args{arr: []string{"NORTH", "EAST", "WEST", "SOUTH", "WEST", "WEST"}}, []string{"WEST", "WEST"}},
		{"5", args{arr: []string{"NORTH", "SOUTH", "SOUTH", "EAST", "WEST", "NORTH", "NORTH"}}, []string{"NORTH"}},
		{"6", args{arr: []string{"EAST", "EAST", "WEST", "NORTH", "WEST", "EAST", "EAST", "SOUTH", "NORTH", "WEST"}}, []string{"EAST", "NORTH"}},
		{"7", args{arr: []string{"NORTH", "EAST", "NORTH", "EAST", "WEST", "WEST", "EAST", "EAST", "WEST", "SOUTH"}}, []string{"NORTH", "EAST"}},
		{"8", args{arr: []string{"WEST", "WEST", "EAST", "EAST", "WEST", "SOUTH", "NORTH", "SOUTH"}}, []string{"WEST", "SOUTH"}},
		{"9", args{arr: []string{"NORTH", "WEST", "EAST", "EAST", "SOUTH"}}, []string{"NORTH", "EAST", "SOUTH"}},
		{"10", args{arr: []string{"SOUTH", "SOUTH", "SOUTH", "NORTH", "SOUTH", "NORTH", "WEST", "EAST"}}, []string{"SOUTH", "SOUTH"}},
		{"11", args{arr: []string{"WEST", "EAST", "EAST", "SOUTH", "NORTH"}}, []string{"EAST"}},
		{"12", args{arr: []string{"WEST", "EAST", "SOUTH", "WEST", "EAST"}}, []string{"SOUTH"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DirReduc(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DirReduc() = %v, want %v", got, tt.want)
			}
		})
	}
}
