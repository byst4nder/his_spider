package main

import "fmt"

func main()  {
	var ptr *int
	fmt.Printf("ptr类型为%T，值为%v \n", ptr, ptr)
	if ptr == nil{
		fmt.Println("空指针")
	}else{
		fmt.Println("非空指针")
	}
	a :=10
	//b :=&a
	var b *int =&a
	fmt.Printf("b的类型%T，数值%v\n",b,b)
	fmt.Println("b的地址", b)
	fmt.Println("*b的值", *b)
	*b++
	fmt.Println("a的值", a)
}
