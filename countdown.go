package sprint

func Countdown(n int) string {
	var vastus string

	if n == 0 {
		return "0!"
	}

	for i := n; i >= 0; i-- {
		if vastus == "" {
			vastus += string('0' + i)
		} else if i == 0 {
			vastus += ", " + string('0'+i) + "!"
		} else {
			vastus += ", " + string('0'+i)
		}

		if i > 0 {
			i--
		}
	}

	if vastus[len(vastus)-1] != '!' {
		vastus += ", 0!"
	}
	return vastus
}
