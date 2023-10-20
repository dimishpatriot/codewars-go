package countipaddress

// https://www.codewars.com/kata/526989a41034285187000de4

import (
	"strconv"
	"strings"
)

func convertToInt(str []string) (res []int) {
	for _, s := range str {
		i, err := strconv.Atoi(s)
		if err == nil {
			res = append(res, i)
		}
	}
	return
}

func IpsBetween(start, end string) int {
	s := convertToInt(strings.Split(start, "."))
	e := convertToInt(strings.Split(end, "."))

	return (e[0]-s[0])<<24 + (e[1]-s[1])<<16 + (e[2]-s[2])<<8 + (e[3] - s[3])
}
