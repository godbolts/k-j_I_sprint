package sprint

func BalanceOut(all []bool) []bool {
	var tõesed int
	var valed int

	for i := 0; i < len(all); i++ {
		if all[i] == true {
			tõesed++
		} else {
			valed++
		}
	}

	diff := tõesed - valed
	if diff > 0 {
		for i := 0; i < diff; i++ {
			all = append(all, false)
		}
	} else if diff < 0 {
		for i := 0; i < -diff; i++ {
			all = append(all, true)
		}
	}

	return all
}
