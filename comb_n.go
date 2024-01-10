package sprint

func CombN(n int) []string {
	if n < 1 || n > 9 {
		return nil
	}

	result := []string{}
	digits := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

	for {
		result = append(result, string(digits[:n]))
		i := n - 1

		for ; i >= 0; i-- {
			if digits[i] < rune('9')-rune(n-i-1) {
				break
			}
		}

		if i < 0 {
			break
		}

		digits[i]++

		for j := i + 1; j < n; j++ {
			digits[j] = digits[j-1] + 1
		}
	}
	return result
}
