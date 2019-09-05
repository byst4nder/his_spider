package main

import "fmt"

// 函数调用本身 有结束条件
// 求阶乘
func main(){
	fmt.Println(getMultiple(10))
	fmt.Println(factorial(10))
}
// 递归的做法实现阶乘
func factorial(n int) int {
	if n ==0{
	return 1
	}
	return n*factorial(n -1)
}
// 通过循环方式实现阶乘
func getMultiple(num int) (result int) {
	result = 1
	for i:=1; i<= num; i++ {
		result *=i
	}
	return
}
/*
使用递归优点：逻辑清晰
使用递归注意事项：防止栈的溢出，函数调用通过栈（stack）这种数据结构实现，
递归调用次数过多，会导致栈的溢出。
*/