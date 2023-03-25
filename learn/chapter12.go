package main

import "fmt"

type Stryc interface {
	Int() int
}

type myCls struct {
	value int
}

func (a *myCls) Int() int {
	return a.value
}
func printTheValueByStruct(arg myCls) {
	fmt.Printf("the value is %d \n", arg.Int())
}
func printTheValue(arg Stryc) int {
	fmt.Printf("the value is %d \n", arg.Int() /*使用接口调用接口的方法*/)
	return arg.Int()
}

func printAnyValue(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println("the type is int")
		case string:
			fmt.Println("the type is string")
		case myCls: /*是自定义数据类型*/
			//使用.(数据类型)进行强转
			var b myCls = arg.(myCls)
			fmt.Println("the type is myCls, the function value is ", b.Int() /*d调用数据类型的方法，golang会转换为数据指针类型调用*/, "the struct value is ", b.value /*调用数据结构的数据*/)
		case Stryc: /*是定义的接口数据类型*/
			fmt.Println("the type is Stryc interface, the function value is ", arg.(Stryc).Int() /*将接口强转到指定接口，并调用方法*/)
		}
	}
}
func main() {
	var V1 *myCls = new(myCls)          //创建一个对象指针
	V1.value = 1111111                  //给对象赋值
	var V2 myCls = myCls{222222222}     //创建一个对象，给对象赋值
	var a interface{} = myCls{33333}    //创建一个对象，将对象赋值后传给一个万能类型接口
	var a1 interface{} = &myCls{444444} //创建一个对象，将对象指针传给一个万能类型接口

	printTheValue(V1)                //V1会转换为Stryc接口被调用其中的方法
	printTheValue(a1.(Stryc))        //万能接口a1中放置的对象指针被强制转为Stryc接口调用
	printTheValueByStruct(*V1)       //强制将V1的对象使用*显示传入函数，因为参数是对象
	printTheValueByStruct(a.(myCls)) //强制将万能接口a中放置的对象转换为对象传入函数，因为参数是对象

	printTheValue(&V2) //将对象的指针传入函数，golang将其转换为Stryc接口
	printAnyValue(V1,  /*传入一个指针，会同Stryc接口数据类型匹配*/
		V2,  /*传入一个对象，会同myCls数据类型匹配*/
		*V1, /*将指针显示为对象传入，会同myCls数据类型匹配*/
		&V2 /*将对象的指针传入，会同Stryc接口数据类型匹配*/)
}
