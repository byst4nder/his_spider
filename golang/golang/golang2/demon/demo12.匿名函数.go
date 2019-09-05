package main

import "fmt"

func main()  {
	//for i:=0;i < 10 ; i++{
	//	fmt.Printf("i = %d \t" , i)
	//	fmt.Println(adder(i))
	//}
	res :=adder()
	fmt.Printf("%T\n", res)
	for i:=0;i<5;i++{
		fmt.Printf("i = %d\t", i)
		fmt.Println(res(i))
	}
}

// 准备实现一个用来计数的函数 ,与环境结合就有记忆功能，（垃圾回收机制）不会当做垃圾被回收
//
func adder( ) func(int) int {
	sum :=0
	res := func(num int) int {
		sum +=num
		return sum
	}
	return res
}