package main

import (
	"fmt"
)

type MyType func (x, y  int) int
//func MyFunc7(a int) int{
//	if a==1{
//		fmt.Println("计算完毕")
//		return 1
//	}
//	return a +MyFunc7(a - 1)
//}
func MyFunc8(c, d int, test11 MyType) (result int) {
	result = test11(c, d)
	return
}

func Add(a, b int) int {
	return a+b
}

func main() {
	//result:=MyFunc7(100)
	//fmt.Println("result=", result)
	//var Test  MyType =MyFunc7
	//result1 :=Test(50)
	//fmt.Println(result1)
   e := MyFunc8(1,3, Add)
   fmt.Println(e)
}
