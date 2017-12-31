package main

func main() {
	a, b, c, x := 1, 2, 3, 2

	switch x {
	case a, b:
		println("a | b")
	case c:
		println("c")
	case 4:
		println("d")
	default:
		println("z")
	}

	// 1. 编译器会确保不会先执行 default 块
	// 2. 相邻的空 case 不构成多条件匹配，
	//    case a: // 隐式转换为： case a: break;
	//    case b:
	//      println("b")
	// 3. 不能出现重复的case常量值
	//    case 1, 2:
	//    case 2: // duplicate case 2 in switch
	// 4. 无需手动执行break，如果需要继续匹配后续条件，则需要执行 fallthrough
	// 5. fallthrough必须放在结尾，break可以终止fallthrough的执行
}
