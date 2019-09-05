package main

import "fmt"

/*
基本数据类型：整形，浮点型，布尔型，字符串，字符
复合数据类型：指针与函数，数组，切片(slice)，map，list，结构体，通道
*/
//声明数组
var arr[]int

func main() {
	a:=[4]float64{2,3,5.4,2.3}
	fmt.Println(a)

	b :=[...]int{1,2,3,4,54,4}
	//遍历数组的形式1
	for i:=0;i <len(b);i++{
		fmt.Print(b[i],"\t")
	}
	fmt.Println()
	// 遍历数组方式二
	for _,value := range b{
		fmt.Print(value,"\t")
	}
	if arr == nil{
		fmt.Println("arr = nil")
	}
	fmt.Print("遍历arr\n")
	// 遍历数组方式三
	for i,value := range b{
		fmt.Print(i,value,"\t")
	}
}