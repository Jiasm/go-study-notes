package main

import (
	"fmt"
)

func test() {

	defer func() {
		fmt.Println(recover())
	}()

	defer func() {
		panic("defer panic1")
	}()

	defer func() {
		panic("defer panic2")
	}()

	panic("test panic")
}

func main() {
	test()
}
