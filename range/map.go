package main

func main() {

	// 直接创建一个map对象
	// 必须直接指定key, value的类型
	userInfo := map[string]string{
		"name": "Niko",
		"age":  "18",
	}

	// range 在处理map时相当于 javascript中的 Object.entries()
	// for (let [key, value] of Object.entries(userInfo))
	for key, value := range userInfo {
		println(key, value)
	}
}
