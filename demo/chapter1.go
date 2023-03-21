package main

import "fmt"

func main() {
	println("hello")
	fmt.Printf("Hello %s", "go")

	var str1 = "Hello"
	println(str1)
	var a, b, c int16 = 1, 2, 3
	fmt.Printf("和: %d \n", a+b+c)
	const (
		a1 = iota
		b1
		c1
		d = "ha"
		e
		e2
		e3
		e4
		f = 100
		g
		h = iota
		i
	)
	fmt.Println(a1, b1, c1, d, e, e2, e3, e4, f, g, h, i)
}
