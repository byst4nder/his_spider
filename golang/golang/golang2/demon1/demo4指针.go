package main

import "fmt"

// 指针存储一个变量的地址，&符号获取地址
func main()  {
	a:=10
	fmt.Printf("变量的地址为%x\n",&a)

	b := []int{1,2,3,4,5,7}
	fmt.Printf("变量的地址为%x\n",&b)
}
/*
指针的特点
指针不能参与运算（不同于C语言），对指针进行运算会报错
*T 指针变量的类型 var指针变量名*指针类型，* 用于指定变量是一个指针，
var ip *int
使用指针  定义指针变量  为指针变量赋值  访问指针变量中只想地址的值。
*/