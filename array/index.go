package main

func main() {
	a := [3]int{1, 2}           // 未初始化元素值为 0
	b := [...]int{1, 2, 3, 4}   // 通过初始化值确定数组长度
	c := [5]int{2: 100, 4: 100} // 使用索引号初始化元素

	d := [...]struct {
		name string
		age  uint8
	}{
		{"user1", 10}, // 可以通过简写的方式，按照struct定义的顺序
		{"user2", 20},
		{
			age:  30,
			name: "user3",
		}, // 也可以通过key: value的方式，但是不要忘记最后的`,`
	}

	println(a[0])
	println(b[1])
	println(c[2])
	println(d[2].age)
}
