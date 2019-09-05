package main

import (
	"fmt"
	"os"
)

func main() {
	m,n :=func(a,b int)(max,min int){
		if a>b{
			max = a
			min = b
		}else {
			max = b
			min = a
		}
		return
	}(30,20)
	fmt.Println("最大值为;",m,"\n","最小值为;",n)
	defer  fmt.Println("aaaaaaaaaaa")
	fmt.Println("bbbbbbbbbbbb")
	defer  fmt.Println("cccccccccc")
	a :=10
	b :=20
	defer func(a,b int){fmt.Println("内部","a = ",a,"b = ",b)}(a, b)
	a = 110
	b = 220
	list1 := os.Args
	n1 := len(list1)
	fmt.Println("n = ", n1)
}
