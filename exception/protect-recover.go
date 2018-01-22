package main

import (
	"fmt"
)

func test(x, y int) {
	var z int

	func() {
		defer func() {
			if recover() != nil {
				z = 0.0
			}
		}()

		z = x / y
		return
	}()

	fmt.Printf("x / y = %d\n", z)
}

func main() {
	test(4, 3)
	test(1, 0)
}
