package isbn10validation

import "testing"

func TestValidISBN10(t *testing.T) {
	type args struct {
		isbn string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"only numbers", args{isbn: "1112223339"}, true},
		{"with X", args{isbn: "048665088X"}, true},
		{"with zeros", args{isbn: "1293000000"}, true},
		{"with small x", args{isbn: "048665088x"}, true},
		{"not valid module", args{isbn: "1234512345"}, false},
		{"short", args{isbn: "1293"}, false},
		{"X first", args{isbn: "X123456788"}, false},
		{"have not numbers", args{isbn: "ABCDEFGHIJ"}, false},
		{"only X's", args{isbn: "XXXXXXXXXX"}, false},
		{"y last", args{isbn: "048665088y"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidISBN10(tt.args.isbn); got != tt.want {
				t.Errorf("ValidISBN10() = %v, want %v", got, tt.want)
			}
		})
	}
}
