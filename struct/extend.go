package main

import (
    "fmt"
)

type user struct {
    name string
    age byte
}

// 声明 所作用于哪个对象 方法的名字 返回值类型 
func (u user) ToString() string {
    return fmt.Sprintf("%+v", u)
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

  println(m.ToString())
}
