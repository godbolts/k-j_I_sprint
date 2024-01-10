package sprint

func FindNextPrime(n int) int {
	if n <= 1 {
		return 2
	}

	prime := n

	found := false
	for !found {

		if IsPrime(prime) {
			found = true
		} else {
			prime += 1
		}
	}
	return prime
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
