package main

import "fmt"

func main(){
	res := Counter()
	fmt.Printf("%T\n", res)
	fmt.Println("res:", res)
	fmt.Println("res:", res())
	fmt.Println("res:", res())
	fmt.Println("res:", res())
	res2 := Counter()
	fmt.Printf("%T\n", res2)
	fmt.Println("res:", res2)
	fmt.Println("res:", res2())
	fmt.Println("res:", res2())
	fmt.Println("res:", res2())
}
//闭包函数，实现计数器功能
func Counter() func() int{
	i:=0
	res := func() int{ //匿名函数
		i++
		return  i
	}
	fmt.Println("Counter内部：",res)
	return res
}