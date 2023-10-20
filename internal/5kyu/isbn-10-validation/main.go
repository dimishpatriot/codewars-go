package isbn10validation

// https://www.codewars.com/kata/51fc12de24a9d8cb0e000001

import "strconv"

func ValidISBN10(isbn string) bool {
	if len(isbn) < 10 {
		return false
	}
	sum := 0
	for i, s := range isbn {
		if i == 9 && (s == 'X' || s == 'x') {
			sum += 10 * 10
		} else {
			n, err := strconv.Atoi(string(s))
			if err != nil {
				return false
			}
			sum += n * (i + 1)
		}
	}
	return sum%11 == 0
}
