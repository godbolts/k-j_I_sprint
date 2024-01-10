package main

import "fmt"

func Abacus(a int, b int) int {
	c := a % b

	return c
}

func main() {
	fmt.Println(Abacus(8, 3))
}
