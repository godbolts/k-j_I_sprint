package sprint

func DigitalRoot(n int) int {
	if len(itoa(n)) == 0 {
		return 0
	}

	vastus := conventeerija(n)
	for len(itoa(vastus)) > 1 {
		vastus = conventeerija(vastus)
	}
	return vastus
}

func conventeerija(arv int) int {
	var summa int
	for arv > 0 {
		arv_j := arv % 10
		summa += arv_j
		arv = arv / 10
	}
	return summa
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}

	var result string
	isNegative := false

	if n < 0 {
		isNegative = true
		n = -n
	}

	for n > 0 {
		digit := n % 10
		result = string('0'+digit) + result
		n = n / 10
	}

	if isNegative {
		result = "-" + result
	}

	return result
}
