package main

type TestStruct struct {
	a *int
}

func fun(ts TestStruct) {
	*ts.a = 1000
}

// Q7: 输出?
func main() {
	intVal := 10

	a := TestStruct{}
	a.a = &intVal
	fun(a)

	println(*a.a)
	// 为何如此: go的穿参和赋值都是值传递, 而struct的值传递则是每个成员变量依次值传递
	//		但是这里我传递的值是一个指针, 就能间接修改了
}
