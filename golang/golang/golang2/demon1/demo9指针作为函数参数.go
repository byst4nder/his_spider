package main

import "fmt"

func main()  {
	a1,b1 :=100, 200
	// 返回值实现数据交换
	a,b := swp0(a1, b1)
	fmt.Println(a,b)
	//使用指针交换
	swap(&a1, &b1)
	fmt.Println(a1, b1)
}
// 具有返回值
func swp0(x, y int)(int, int){
	return y, x
}
// 使用指针作为参数的写法
func swap(x, y *int)  {
	*x, *y = *y, *x
}