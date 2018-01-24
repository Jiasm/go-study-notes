package main

import (
	"fmt"
)

func main() {
	s1 := []int{0, 1, 2, 3, 8: 100}
	fmt.Println(s1, len(s1), cap(s1))

	s2 := make([]int, 6, 8)
	fmt.Println(s2, len(s2), cap(s2))

	s3 := make([]int, 6)
	fmt.Println(s3, len(s3), cap(s3))
}
