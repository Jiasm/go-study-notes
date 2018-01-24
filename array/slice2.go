package main

import (
	"fmt"
)

func main() {
	s := []int{0, 1, 2, 3}

	p1 := &s[2]
	*p1 += 100 // 会影响到数组

	p2 := s[3]
	p2 += 200 // 不会影响到数组

	s2 := s[:]
	s2[1] += 50 // 会影响到数组

	fmt.Println(s)
}
