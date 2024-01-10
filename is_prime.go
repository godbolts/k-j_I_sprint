package sprint

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
