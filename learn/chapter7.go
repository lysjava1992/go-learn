package main

func main() {
	var a int64 = 15
	var ip *int64
	ip = &a

	println(&a)
	println(ip)
	println(*ip)

	var ptr *int64
	println(ptr)
	println(nil == ptr)
}
