package main

// 闭包
func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}
func main() {
	nextNumber := getSequence()
	println(nextNumber())
	println(nextNumber())
	println(nextNumber())
	println(nextNumber())
	nextNumber1 := getSequence()
	println(nextNumber1())
	println(nextNumber1())

}
