package main

import (
  "fmt"
)

type user struct {
  name string
  age byte
}

type manager struct {
  user
  title string
}

func main() {
  var m manager

  m.name = "Niko"
  m.age = 29
  m.title = "CTO"

  fmt.Println(m)
}
