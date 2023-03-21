package main

import "fmt"

func main() {
	var a, b int = 100, 2023
	fmt.Printf("swap0 a :%d -- b: %d \n", a, b)
	swap1(a, b)
	fmt.Printf("swap0 a :%d -- b: %d \n", a, b)
	swap2(&a, &b)
	fmt.Printf("swap0 a :%d -- b: %d \n", a, b)
}

// 值传递
func swap1(a int, b int) {
	temp := a
	a = b
	b = temp
	fmt.Printf("swap1 a :%d -- b: %d \n", a, b)
}

// 引用传递
func swap2(a *int, b *int) {

	temp := *a
	*a = *b
	*b = temp
	fmt.Printf("swap2 a :%d -- b: %d \n", a, b)
}
