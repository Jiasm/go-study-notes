package main

type Person struct {
	name string
	age  uint8
}

func main() {
	personChan := make(chan Person, 2)

	p1 := Person{
		name: "Niko",
		age:  18,
	}

	p2 := Person{
		name: "Roman",
		age:  20,
	}

	personChan <- p1
	personChan <- p2

	close(personChan)

	for chanItem := range personChan {
		println(chanItem.name, chanItem.age)
	}
}
