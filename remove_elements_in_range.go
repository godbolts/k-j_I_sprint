package sprint

func RemoveElementsInRange(all []float64, from, to int) []float64 {
	if to < from {
		from, to = to, from
	}

	if from < 0 {
		from = 0
	}
	if to < 0 {
		to = 0
	}
	if from > len(all) {
		from = len(all)
	}
	if to >= len(all) {
		to = len(all)
	}

	kokku := append(all[:from], all[to:]...)

	return kokku
}
