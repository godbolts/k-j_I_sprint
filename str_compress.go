package sprint

func StrCompress(input string) string {
	if len(input) == 0 {
		return ""
	}

	var vastus string
	count := 1

	for i := 1; i < len(input); i++ {
		if input[i] == input[i-1] {
			count++
		} else {
			if count > 1 {
				vastus += Itoa(count)
			}
			vastus += string(input[i-1])
			count = 1
		}
	}

	if count > 1 {
		vastus += Itoa(count)
	}
	vastus += string(input[len(input)-1])

	return vastus
}

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}

	var result string

	for n > 0 {
		digit := n % 10
		result = string('0'+digit) + result
		n /= 10
	}

	return result
}
