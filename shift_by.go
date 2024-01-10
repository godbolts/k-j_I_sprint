package sprint

func ShiftBy(r rune, step int) rune {
	if step > 26 || step < -26 {
		step = step % 26
	}

	var a rune

	switch {
	case r+rune(step) > 122:
		a = rune(int(r)+step) - 26
	case r+rune(step) < 97:
		a = rune(int(r)+step) + 26
	default:
		a = rune(int(r) + step)
	}

	return a
}
