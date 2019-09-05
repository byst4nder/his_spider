package main

import (
	"fmt"
)

func main() {
	a, b:=20, 10
	// 一行一行执行，自动换行
	//fmt.Println("a =",a, a)
	//fmt.Println("a =",a, b)
	a, b=b,a // 交换两个变量的值
	fmt.Printf("a = %d\n",b)
	fmt.Printf("a = %d\n",a)
	// 匿名变量 ，丢弃数据不处理
	var tmp int
	tmp,_ = a,b
	fmt.Println(tmp)
}