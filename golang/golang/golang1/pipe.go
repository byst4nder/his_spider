package main

import "fmt"
/*
*1.简单的channel队列现进先出特性
*在主线程上建立单个协程，完成某项工作后，关闭该管道,关闭与否为可选项。
*停止一段阻塞有两种方式，chanInt内能读取到值，即非空;chanInt被关闭
*/

func main()  {
	pipe := make(chan int,10)  // 创建管道  make channel 类型的管道  存储int类型  三个
	pipe <- 1 //向管道里添加参数
	pipe <- 2
	pipe <- 3
	var t1 int
	t2 :=<- pipe
	var t3 int
	var t4 =<- pipe

	t1 =<- pipe // 从pipe里取数据
	fmt.Println(t1)
	fmt.Println(t2, t3, t4)
	fmt.Println(len(pipe))
	fmt.Println(pipe)  // 打印出来的是内存地址
}