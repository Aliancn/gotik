package main

// 那么go提供的方法是怎么一个回事呢? 为什么有的时候这么写
/*
func (o Obj) f()
*/

// 有的时候又这么写呢?
/*
func (o *Obj) f()
*/

// go没有面向对象, 上面仅仅是面向过程的语法糖, go也没有多态继承等面向对象的特征
// 	下面我将验证

type Object struct {
	testInt int
}

func (o *Object) Do() {
	println(o)
}

func main() {
	o := Object{}
	o.Do()

	var o2 *Object = nil
	o2.Do()

	// 结论, go的面向对象写法只是传参的语法糖, 让它看起来是面向对象而已, 我认为可能编译器是大概这样做的:

	/*
		生成Do(*Object)的函数
		调用o.Do()的时候其实就是间接调用Do(o)罢了
	*/

	// o.Do的写法叫选择子, 可以把对象绑定到方法上, 看下面的例子
	f := o.Do
	f()
}
