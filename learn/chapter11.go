package main

import (
	"fmt"
)

type Writer interface {
	Writer([]byte) (int, error)
}
type StringWriter struct {
	str string
}

func (sw *StringWriter) Writer(data []byte) (int, error) {
	sw.str += string(data)
	return len(data), nil
}
func main() {
	var i interface{} = "hello,world"
	// x.(T) 类型断言用于提取接口的基础值
	// x接口类型 T要断言的类型
	//使用类型断言将 i 转换为字符串类型
	str, ok := i.(string)
	if ok {
		fmt.Printf(" '%s' is a string \n", str)
	} else {
		fmt.Println("转换失败")
	}

	var w Writer = &StringWriter{}
	//T(x)
	// 结果类型:=<目标类型>（<表达式>）
	//类型转换是用来在不同但相互兼容的类型之间的相互转换的方式，所以，当类型不兼容的时候，是无法转换的
	//使用类型转换将 w 转换为 StringWriter 类型，并将转换后的值赋值给变量 sw
	sw := w.(*StringWriter)
	sw.str = "Hello,World"
	fmt.Println(sw.str)
}
