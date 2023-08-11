package main

func SetVal(v *int) {
	if v == nil {
		panic("unexpected")
	}

	*v = 100
}

// 指针的妙用, 间接传参
func main() {
	// Q4: 结果?
	a := 10
	SetVal(&a)
	println(a)

	// 为何如此? 因为&a是a的地址, 我们把这个值传递给了*int变量, 对方就能间接地改变
	//		"变量"本体
}
