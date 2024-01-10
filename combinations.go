package sprint

import "fmt"

func Combinations() string {
	var vastus string
	for i := 0; i <= 9; i++ {
		for j := i + 1; j <= 9; j++ {
			for z := j + 1; z <= 9; z++ {
				if i != j && i != z && j != z && i < j && j < z {
					if vastus != "" {
						vastus += ", "
					}
					vastus += fmt.Sprintf("%d%d%d", i, j, z)
				}
			}
		}
	}

	return vastus //
}
