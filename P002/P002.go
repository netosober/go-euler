package main

import "fmt"

func main() {

	a, b, sum := 1, 1, 0

	for a < 4000000 {
		fmt.Printf("%d\n", a)
		c := a
		a = a + b
    b = c
		if a%2 == 0 {
			sum += a
		}
	}
	fmt.Printf("%d\n", sum)

}
