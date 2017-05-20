package main

func main() {
    x := 100
    p := (*int)(&x) // 此处 *int的括号不能省去，否则会被解析为： *(int(&x))

    println(p)
}
