package main

func main() {
	// Q3: 下面会有输出吗?
	var m = make(map[string]int)

	var m2 = m

	m2["hello"] = 1

	for _, v := range m {
		println(v)
	}

	// 结论: 引用传递是指针的"值传递", java里只要引用, 其实更类似为指针的值传递
	// 		比如String a = "你好"; String b = a; b = "不好"; print(a)发现结果是"你好"
}
