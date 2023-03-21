package main

import "math"

type cb func(float64) float64

func main() {
	testCallBack(81, callBack)

}
func testCallBack(x float64, f cb) {
	println(f(x))
}
func callBack(x float64) float64 {
	return math.Sqrt(x)
}
func callBack2(x string) (string, string) {
	return x, x + "44"
}
