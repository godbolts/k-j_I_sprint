package sprint

func Doop(a int, op string, b int) int {
	var vastus int
	switch op {
	case "+":
		vastus = a + b
	case "-":
		vastus = a - b
	case "/":
		if b == 0 {
			vastus = 0
		} else {
			vastus = a / b
		}
	case "*":
		vastus = a * b
	case "%":
		if b == 0 {
			vastus = 0
		} else {
			vastus = a % b
		}
	default:
		vastus = 0
	}

	return vastus

}
