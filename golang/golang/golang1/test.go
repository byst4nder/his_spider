package main

import "fmt"

// 定义add 函数  指定a，b 参数的数据类型
//func add(a int, b int) int {
//	var sum int  // 定义变量指定变量类型
//	sum = a + b   //计算加法
//	return sum //返回计算结果
//}
// 定义main函数

	//var c int  //定义c变量
	//c = add(100, 200)  //调用函数将函数赋值给c
	//go test_goroute(300,300)
	//fmt.Println("add(100,200)=",c)
	//for i:=0; i<100; i++{
	//
	//}

/*
1 天然并发
从语言层面支持并发，非常简单
goroute，轻量级线程， 创建成千上万个goroute成为可能
基于CSP模型
 */
// 相加
func add(a int, b int) int {
	sum := a + b
	return sum

}
// 循环
func aa() {
	for i:=0; i<10; i++{
		println(i)
	}
}

func main() {
	c :=add(2, 3)
	fmt.Println("2 + 3 =", c)
	aa()
}

