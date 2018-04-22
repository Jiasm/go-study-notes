package main

func main() {
	// 创建一个数组，必须指定类型，长度可以省略
	data := [3]string{"a", "b", "c"}

	// 写了...表示为数组
	// 没有写则表示为切片
	data2 := [...]string{"a", "b", "c", "d"}

	data3 := append(data2, "e")

	for i, item := range data {
		println(i, item)
	}

	for i, item := range data3 {
		println(i, item)
	}
}
