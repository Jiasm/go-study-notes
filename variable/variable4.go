package main

var i int = 1

func getCount() int {
    i++
    return i
}

func main() {
  x, y := getCount(), getCount()

  println("x: %d, y: %d", x, y)
}
