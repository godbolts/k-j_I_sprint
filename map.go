package sprint

func Map(f func(int) bool, a []int) []bool {
	var vastus []bool

	for _, element := range a {
		if f(element) {
			vastus = append(vastus, true)
		} else {
			vastus = append(vastus, false)
		}
	}
	return vastus
}

func IsNegative(n int) bool {
	return n < 0
}

func IsPrime(n int) bool {
	on := true
	if n <= 1 {
		on = false
	}

	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			on = false
		}
	}
	return on
}
