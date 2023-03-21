package main

import (
	"fmt"
	"math"
)

type Vertex2 struct {
	X, Y float64
}

// Abs 为类型Vertex2声明一个方法
// 基于值传递是对v值副本的操作
func (v Vertex2) Abs() float64 {
	v.X = v.X * 100
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Scale Scale为类型Vertex2引用声明的一个方法
// 引用传递会修改v的原始值
func (v *Vertex2) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
func main() {
	v := Vertex2{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(v)
	v.Scale(10)
	fmt.Println(v)
}
