package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
	N()
}

type T struct {
	S string
}

// N 类型通过实现一个接口的所有方法来实现该接口。既然无需专门显式声明
func (t T) N() {
	panic("implement me")
}

// M 类型通过实现一个接口的所有方法来实现该接口。既然无需专门显式声明
func (t T) M() {
	t.S = t.S + "100"
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}
func (f F) N() {
	fmt.Println(f)
}

// 接口也是值可以传递
func describe(i I) {
	i.M()
	fmt.Printf("(%v, %T)\n", i, i)
}
func main() {
	// 当 T 必须实现了 M 和 N（所有）的接口方法 T才是I的实现
	var i I = &T{"hello"}
	describe(i)
	//i.M()
	i = F(math.Pi)
	describe(i)
	//i.M()

	// 空接口可以保存任何类型的值
	// 因为每个类型都至少实现了零个方法
	var g interface{}
	describeG(g)
	g = 42
	describeG(g)
	g = "king"
	describeG(g)
}

func describeG(g interface{}) {
	fmt.Printf("(%v, %T)\n", g, g)
}
