package main

import "fmt"

type Vertex3 struct {
	X, Y float64
}

func (v *Vertex3) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex3, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
func main() {
	v := Vertex3{3, 4}
	//Scale的接受者为指针类型
	//此时 v.Scale(10)会解释为(&v).Scale(10)
	// 等同于p.Scale(10)
	v.Scale(10)
	fmt.Println(v)
	p := &v
	p.Scale(10)
	fmt.Println(v)
	ScaleFunc(&v, 10)
	//函数ScaleFunc参数类型为指针 则非指针类型编译不过
	//ScaleFunc(v, 10)
	fmt.Println(v)
}
