package main

func main() {
  x := 10
  y := 20

  switch {
    case y < 0:
      println("-y")
    case x > 0:
      println("x")
  }
}
