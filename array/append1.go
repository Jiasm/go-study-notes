package main

import (
	"fmt"
)

func main() {
	s := make([]int, 0, 5)
	fmt.Printf("%p\n", &s)

	s2 := append(s, 1) // 如果超过原数组上限，则会重新在内存中创建一个数组
	fmt.Printf("%p\n", &s2)

	fmt.Println(s, s2, cap(s), cap(s2))
}
