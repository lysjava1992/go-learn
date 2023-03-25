package main

import (
	"fmt"
	"math"
)

// Abser 接口类型 ：由一组方法签名定义的集合
// 接口类型的变量 可以保存任何实现了这些方法的（值）
type Abser interface {
	Abs() float64
}

type MyFloat float64

// Abs MyFloat 实现了 Abser
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex5 struct {
	X, Y float64
}

// Abs *Vertex 实现了 Abser
func (v *Vertex5) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	var a Abser
	f := MyFloat(-56)
	v := Vertex5{3, 4}
	a = f
	fmt.Println(a.Abs())
	a = &v
	fmt.Println(a.Abs())
	// 下面一行，v 是一个 Vertex5（而不是 *Vertex5）
	// 所以没有实现 Abser。
	//a = v

}
