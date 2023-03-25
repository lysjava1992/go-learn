package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	fmt.Println(math.Pi)
	switch {
	case time.Now().Hour() < 12:
		println("------morning--")
	case time.Now().Hour() < 17:
		println("------afternoon--")
	default:
		println("------evening------")
	}
	// defer 先进后出
	defer fmt.Printf("--4--")
	defer fmt.Printf("--3--")
	defer fmt.Printf("world!")
	fmt.Printf("hello")
	fmt.Printf("--1--")
	fmt.Printf("--2--")
}
