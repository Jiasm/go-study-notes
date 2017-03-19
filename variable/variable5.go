package main

func getValue() (int, string) {
    return 1, "Niko"
}

func main() {
    age, name := getValue()

    println("age: ", age, "name: ", name)
}
