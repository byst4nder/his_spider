package main

import (
	"fmt"
)

const count int =4
func main(){
	a:=[count]string{"abc","ABC","123","一二三"}
	// 查看数组的指针的类型和值
	fmt.Printf("%T, %v \n", &a, &a)
	// 定义指针数组  和数组指针的区别
	var ptr [count]*string
	fmt.Printf("%T, %v \n", ptr, ptr)
	for i:=0; i< count; i++ {
		//将数组中每个元素的地址一次赋值给指针的每个元素
		ptr[i] = &a[i]
	}
	fmt.Printf("%T , %v \n", ptr, ptr)
	//根据指针数组的每个地址获取该地址锁只想的元素的真实数值
	for i:=0; i<count;i++ {
		fmt.Println(*ptr[i])
	}
	for _,value :=range ptr{
		fmt.Println(*value)
	}
}
/*
指针的指针 var ptr **int
*/
