package main

import "fmt"

func main(){
	var a int
	var ptr *int
	var pptr **int

	a = 123
	// 为指针地址赋值
	ptr = &a
	fmt.Println("ptr:", ptr)
	//为批pptr 赋值
	pptr = &ptr
	fmt.Println("PPtr:",pptr)
	fmt.Println("&PPtr:",&pptr)
	// 获取指针对应的值
	fmt.Printf("变量a = %d\n", a)
	fmt.Printf("指针的变量*ptr = %d\n", *ptr)
	fmt.Printf("指向到指针的变量**pptr = %d\n", **pptr)
}


















