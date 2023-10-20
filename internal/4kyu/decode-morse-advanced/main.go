package decodemorseadvanced

// https://www.codewars.com/kata/54b72c16cd7f5154e9000457

import (
	"strings"

	"github.com/dimishpatriot/codewars-go/internal/6kyu/decode-morse" // remove from solution!
)

var DecodeMorse = decodemorse.DecodeMorse // remove from solution

func findSeq(str string, startKey, stopKey byte) []int {
	res := []int{}

	for i := 0; i < len(str); i++ {
		if str[i] == startKey {
			for j := i + 1; j < len(str); j++ {
				if str[j] == stopKey {
					res = append(res, j-i)
					i = j + 1
					break
				}
			}
		}
	}
	return res
}

func min(seq []int) int {
	if len(seq) == 0 {
		return 0
	}
	res := seq[0]
	for _, p := range seq {
		if p < res {
			res = p
		}
	}
	return res
}

func DecodeBits(bits string) string {
	result := []string{}
	clearBits := strings.Trim(bits, "0")
	pauses := findSeq(clearBits, '0', '1')
	signals := findSeq(clearBits, '1', '0')
	signal := min(signals)

	speed := len(clearBits)
	if len(pauses) > 0 {
		pause := min(pauses)
		switch {
		case signal <= pause:
			speed = signal
		case pause%7 == 0:
			speed = pause / 7
		case pause%3 == 0:
			speed = pause / 3
		default:
			speed = pause
		}
	}

	wordsSeq := strings.Split(clearBits, strings.Repeat("0", speed*7))
	for _, word := range wordsSeq {
		wordChars := strings.Split(word, strings.Repeat("0", speed*3))
		encodedWord := []string{}

		for _, charSeq := range wordChars {
			chars := strings.Split(charSeq, strings.Repeat("0", speed))
			dashDots := []rune{}

			for _, c := range chars {
				switch c {
				case strings.Repeat("1", speed*1):
					dashDots = append(dashDots, rune(46))
				case strings.Repeat("1", speed*3):
					dashDots = append(dashDots, rune(45))
				}
			}

			encodedWord = append(encodedWord, string(dashDots))
		}
		result = append(result, strings.Join(encodedWord, " "))
	}

	return strings.Join(result, "   ")
}
