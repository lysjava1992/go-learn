package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	// p是一个指针
	// p是一个string类型变量的指针
	var p *string
	a := "张三"
	// p是a的指针
	p = &a
	println(p)
	//*p即指针所指向的变量a的值
	println(*p)
	*p = "king"
	println(a)

	v := Vertex{1, 2}

	h := &v
	v.X = v.X * 10
	h.Y = h.Y * 10
	fmt.Println(v)

	o := Vertex{3, 4}
	g := Vertex{3, 4}
	O := &o
	G := &g
	fmt.Println(o)
	fmt.Println(g)
	fmt.Println(o == g)   //true
	fmt.Println(O == G)   //false
	fmt.Println(*O == *G) //true

	var s []int
	fmt.Println(s, len(s), cap(s))
	fmt.Println(s == nil)

}
