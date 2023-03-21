package main

import "fmt"

type Hello interface {
	Hello()
}
type Say struct {
	msg string
}

func (s Say) Hello() {
	fmt.Println(s.msg)
}

type Say2 struct {
	a int
}

func (s Say2) Hello() {
	fmt.Println(s.a)
}

/*
类型断言

	value:ok=i.(T)
	 判断一个接口值 i 是否保存了特定的类型
*/
func main() {
	var h Hello = Say{"Hello Go"}
	h.Hello()
	s := h.(Say)
	fmt.Println(s)
	f, ok := h.(Say2)
	fmt.Println(f, ok)
	f2, ok2 := h.(Say)
	fmt.Println(f2, ok2)
	do("15")
}

// 类型选择
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)

	}
}
