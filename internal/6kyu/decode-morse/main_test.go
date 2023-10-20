package decodemorse

import "testing"

func TestDecodeMorse(t *testing.T) {
	type args struct {
		morseCode string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"basic", args{morseCode: ".-"}, "A"},
		{"AZ", args{morseCode: ".- .-.-.- .-.-.- --.."}, "A..Z"},
		{"basic", args{morseCode: "."}, "E"},
		{"basic", args{morseCode: ".."}, "I"},
		{"basic", args{morseCode: ". ."}, "EE"},
		{"basic", args{morseCode: ".   ."}, "E E"},
		{"basic", args{morseCode: "..."}, "S"},
		{"basic", args{morseCode: "---"}, "O"},
		{"SOS 1", args{morseCode: "...---..."}, "SOS"},
		{"SOS 2", args{morseCode: "... --- ..."}, "SOS"},
		{"SOS 3", args{morseCode: "...   ---   ..."}, "S O S"},
		{"complex", args{morseCode: " . "}, "E"},
		{"complex", args{morseCode: "   .   . "}, "E E"},
		{"complex", args{morseCode: "      ...---... -.-.--   - .... .   --.- ..- .. -.-. -.-   -... .-. --- .-- -.   ..-. --- -..-   .--- ..- -- .--. ...   --- ...- . .-.   - .... .   .-.. .- --.. -.--   -.. --- --. .-.-.-  "}, "SOS! THE QUICK BROWN FOX JUMPS OVER THE LAZY DOG."},
		{"site", args{morseCode: ".... - - .--. ... ---... -..-. -..-. --. --- --- --. .-.. . .-.-.- -.-. --- --"}, "HTTPS://GOOGLE.COM"},
		{"email", args{morseCode: ". -..- .- -- .--. .-.. . .--.-. - . ... - .-.-.- -.-. --- --"}, "EXAMPLE@TEST.COM"},
		{"money", args{morseCode: ".---- ..--.- ----- ----- ----- ..--.- ----- ----- ----- ..--.- ----- ----- ----- / ...-..-"}, "1_000_000_000$"},
		{"numbers", args{morseCode: "----- .---- ..--- ...-- ....- ..... -.... --... ---.. ----."}, "0123456789"},
		{"math", args{morseCode: ".---- .-.-. ..--- -...- ...--"}, "1+2=3"},
		{"say", args{morseCode: "... .- -.-- ---...     .-..-. .... . .-.. .-.. --- --..--     .-- --- .-. .-.. -.. -.-.-- .-..-.     .-...     .--- --- -.--"}, "SAY: \"HELLO, WORLD!\" & JOY"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DecodeMorse(tt.args.morseCode); got != tt.want {
				t.Errorf("DecodeMorse() = %v, want %v", got, tt.want)
			}
		})
	}
}
