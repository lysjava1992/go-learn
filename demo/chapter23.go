package main

import "fmt"

type Vertex4 struct {
	X, Y float64
}

func (v Vertex4) Abs(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func AbsFunc(v Vertex4, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

/*
使用指针作为接收者：

	1.方法能够修改其接收者的值
	2.避免方法调用时复制该值，更高效
*/
func main() {
	v := Vertex4{3, 4}
	//Abs的接受者为值

	v.Abs(10)
	fmt.Println(v) //{3,4}
	p := &v
	//此时 p.Abs(10)会解释为(*p).Abs
	// 等同于v.Abs(10)
	//是值传递
	p.Abs(10)
	fmt.Println(v) //{3,4}
	//函数ScaleFunc参数类型为值 则指针类型编译不过
	//AbsFunc(&v, 10)
	AbsFunc(v, 10)
	fmt.Println(v) //{3,4}
}
