package sumofdigits

import "testing"

func TestDigitalRoot(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		args    args
		wantSum int
	}{
		{"", args{n: 16}, 7},
		{"", args{n: 195}, 6},
		{"", args{n: 992}, 2},
		{"", args{n: 167346}, 9},
		{"", args{n: 0}, 0},
		{"", args{n: 10}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSum := DigitalRoot(tt.args.n); gotSum != tt.wantSum {
				t.Errorf("DigitalRoot() = %v, want %v", gotSum, tt.wantSum)
			}
		})
	}
}
