package main

import "strconv"

func main() {
	var a int16 = 12
	var b int16 = 5
	c := float32(a) / float32(b)
	println(c)
	var str1 = "15"
	n, _ := strconv.Atoi(str1)
	println(n + 1)
	println("----" + strconv.Itoa(15))

}
