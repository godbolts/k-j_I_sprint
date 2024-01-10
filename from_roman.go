package sprint

var m = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func FromRoman(s string) int {
	if s == "" {
		return 0
	}
	is := []rune(s)
	var c0 rune
	var cv0 int
	var r int // Define r here

	for i := len(is) - 1; i >= 0; i-- {

		c := is[i]
		k := c == '\u0305'
		if k {
			if i == 0 {
				return 0
			}
			i--
			c = is[i]
		}
		cv := m[c]
		if cv == 0 {
			if c == 0x0305 {
				return 0
			} else {
				return 0
			}
		}
		if k {
			c = -c
			cv *= 1000
		}

		switch {
		default:
			fallthrough
		case c0 == 0:
			c0 = c
			cv0 = cv
		case c == c0:
		case cv*5 == cv0 || cv*10 == cv0:

			c0 = 0
			r -= cv
			continue
		}
		r += cv
	}
	return r
}
