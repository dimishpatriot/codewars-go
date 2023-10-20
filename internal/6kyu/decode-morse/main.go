package decodemorse

// https://www.codewars.com/kata/54b724efac3d5402db00065e

import (
	"strings"

	"github.com/dimishpatriot/codewars-go/pkg/morse" // remove from solution!
)

var MORSE_CODE = morse.EN // remove from solution

func DecodeMorse(morseCode string) string {
	result := []string{}
	words := strings.Split(strings.Trim(morseCode, " "), "   ")

	for _, w := range words {
		chars := strings.Split(w, " ")
		decoded := ""

		for _, c := range chars {
			decoded += MORSE_CODE[c]
		}

		result = append(result, decoded)
	}

	return strings.Join(result, " ")
}
