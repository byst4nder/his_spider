package main

import "fmt"

func main()  {
	a:=10
	fmt.Println("函数调用前a的值",a)
	b :=&a
	change(b)
	fmt.Println("函数调用之后的a值",a)
	change_1(a)
	fmt.Println("第二次修改a的值后", a)
}
func change(num *int){
	*num = 20
}

func change_1(num int)  {
	num = 30

}