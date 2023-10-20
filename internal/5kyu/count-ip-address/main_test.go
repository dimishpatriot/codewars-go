package countipaddress

import "testing"

func TestIpsBetween(t *testing.T) {
	type args struct {
		start string
		end   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{start: "10.0.0.0", end: "10.0.0.50"}, 50},
		{"", args{start: "20.0.0.10", end: "20.0.1.0"}, 246},
		{"zero", args{start: "150.0.0.0", end: "150.0.0.0"}, 0},
		{"one", args{start: "150.0.0.0", end: "150.0.0.1"}, 1},
		{"", args{start: "10.11.12.13", end: "10.11.13.0"}, 243},
		{"2^8", args{start: "160.0.0.0", end: "160.0.1.0"}, 256},
		{"2^16", args{start: "170.0.0.0", end: "170.1.0.0"}, 65536},
		{"", args{start: "50.0.0.0", end: "50.1.1.1"}, 65793},
		{"2^24", args{start: "180.0.0.0", end: "181.0.0.0"}, 16777216},
		{"nice", args{start: "1.2.3.4", end: "5.6.7.8"}, 67372036},
		{"full 2^32", args{start: "0.0.0.0", end: "255.255.255.255"}, 4294967295},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IpsBetween(tt.args.start, tt.args.end); got != tt.want {
				t.Errorf("IpsBetween() = %v, want %v", got, tt.want)
			}
		})
	}
}
