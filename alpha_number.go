package sprint

func AlphaNumber(n int) string {
	if n == 0 {
		return "a"
	}
	var result string
	var neg bool
	if n < 0 {
		n = -n
		neg = true
	}
	for n > 0 {
		digit := n % 10
		result = string('a'+digit) + result
		n /= 10
	}
	if neg {
		return "-" + result
	}
	return result //
}
