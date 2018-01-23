package main

import (
	"fmt"
)

/**
 * copy操作会直接作用在第一个参数上，修改原始数据
 */
func main() {
	arr := [3]int{1: 1, 2: 2}

	println(len(arr), cap(arr))

	sli1 := []int{1, 2, 3, 4, 5}
	t1 := []int{11, 21, 31, 41, 5, 6, 7, 8, 9}

	copy(sli1, t1) // [11 21 31 41 5]

	fmt.Println(sli1)

	sli2 := []int{1, 2, 3, 4, 5}
	t2 := []int{11, 21, 31, 41, 5, 6, 7, 8, 9}

	copy(t2, sli2) // [1 2 3 4 5 6 7 8 9]

	fmt.Println(t2)
}
