package directionsreduction

// https://www.codewars.com/kata/550f22f4d758534c1100025a

func DirReduc(arr []string) []string {
	if len(arr) < 2 {
		return arr
	}

	patient := make([]string, len(arr))
	copy(patient, arr)

	for {
		leftSide := make([]string, 0)
		startLen := len(patient)

		for i := 0; i < len(patient)-1; i += 1 {
			pair := [2]string{patient[i], patient[i+1]}
			if havePairs(pair) {
				rightSide := patient[i+2:]
				patient = append(leftSide, rightSide...)
				break
			} else {
				leftSide = append(leftSide, patient[i])
			}
		}

		if len(patient) == startLen {
			break
		}
	}
	return patient
}

var pairs = [][2]string{
	{"NORTH", "SOUTH"},
	{"SOUTH", "NORTH"},
	{"WEST", "EAST"},
	{"EAST", "WEST"},
}

func havePairs(currentPair [2]string) bool {
	for _, p := range pairs {
		if p == currentPair {
			return true
		}
	}
	return false
}
