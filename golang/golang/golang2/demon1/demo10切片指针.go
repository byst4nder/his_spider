package main

import "fmt"

// 修改切片的值
func main(){
	a:= []int{1,2,3,4}
	modify(&a)
	fmt.Println(a)
}
func modify(arr *[]int){
	(*arr)[0] = 250
}
/*
指针数组
*/