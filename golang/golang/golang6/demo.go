package main

import "fmt"

func main()  {
	// 变量：程序运行期间，可以改变的量，var声明
	// 常量：运行期间不可以改变，const声明
	const a, c int = 10, 20
	const b, e= 20.0, 40
	var g = 34
	fmt.Println(a,  b, c, e)
	fmt.Printf("%T\n", b)
	fmt.Printf("g的类型为：%T", g)
}
