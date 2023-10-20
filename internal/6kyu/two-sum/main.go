package twosum

// https://www.codewars.com/kata/52c31f8e6605bcc646000082

func TwoSum(numbers []int, target int) [2]int {
	result := [2]int{}

	for i, n := range numbers {
		rest := target - n
		if rest > 0 {
			result[0] = i

			for j, m := range numbers[i+1:] {
				if rest == m {
					result[1] = j + i + 1
					return result
				}
			}

		}
	}
	return result
}
