package sumofdigits

// https://www.codewars.com/kata/541c8630095125aba6000c00

func DigitalRoot(n int) (sum int) {
	if n < 10 {
		return n
	}

	for n > 0 {
		rest := n % 10
		n = (n - rest) / 10
		sum += rest
	}

	if sum >= 10 {
		sum = DigitalRoot(sum)
	}

	return sum
}
