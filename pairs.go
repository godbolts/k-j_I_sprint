package sprint

import "fmt"

func Pairs() string {
	var vastus string
	for i := 0; i <= 99; i++ {
		for j := 0; j <= 99; j++ {
			if i < j {
				vastus += fmt.Sprintf("%02d %02d, ", i, +j)
			}
		}
	}
	if len(vastus) > 2 {
		vastus = vastus[:len(vastus)-2]
	}
	return vastus
}
