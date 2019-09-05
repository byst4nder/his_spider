package main

import (
	"fmt"
	"math"
)

func main(){
	// 定义时调用
	func(data int){
		fmt.Println("你的成绩:",data)
	}(98)
	// 在定义时调用有参数匿名函数
	result := func (data float64) float64 {
		return math.Sqrt(data)
	}(250)
	fmt.Println("平方根：",result)
	// 将匿名函数赋值给变量，需要时再调用
	myfunc := func(data string) string{
		return data
	}
	fmt.Println(myfunc("欢迎学习Go语言"))
}
