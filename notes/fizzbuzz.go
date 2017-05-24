package main

import "fmt"

func main() {
    for number := 1; number <= 100; number++ {
        if number % 15 == 0 {
	    fmt.Printf("number: %d, with FizzBuzz\n", number)
	} else if number % 3 == 0 {
	    fmt.Printf("number: %d, with Fizz\n", number)
	} else if number % 5 == 0 {
	    fmt.Printf("number: %d, with Buzz\n", number)
	} else {
	    fmt.Printf("number: %d\n", number)
        }
    }
}
