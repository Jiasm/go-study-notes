package main

/**
 * panic    用来抛出异常
 * recover  用来捕获异常
 */

func test() {
	defer func() {
		if err := recover(); err != nil {
			println(err.(string))
		}
	}()

	panic("panic error!")
}

func main() {
	test()
}
