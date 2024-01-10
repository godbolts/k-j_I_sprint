package sprint

func Payout(amount int, denominations []int) (payout []int) {
	// Sort the denominations in decreasing order.
	for i := 0; i < len(denominations); i++ {
		for j := i + 1; j < len(denominations); j++ {
			if denominations[i] < denominations[j] {
				denominations[i], denominations[j] = denominations[j], denominations[i]
			}
		}
	}

	for _, denom := range denominations {
		for amount >= denom {
			amount -= denom
			payout = append(payout, denom)
		}
	}

	// If the amount is not zero, it means the payout cannot be made.
	if amount != 0 {
		return []int{}
	}

	return payout
}
