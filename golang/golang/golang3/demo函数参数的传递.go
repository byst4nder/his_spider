package main

import "fmt"

func main(){
	a :=10
	fmt.Printf("变量a的内存地址：%p， 值为：%v\n" , &a, a)
	changeIntVal(a)
	changeInPtr(&a)
}
func changeIntVal(a int){
	fmt.Printf("---------cahngeIntVal函数内存\n")
	fmt.Printf("-----a的内存地址为：%p， 值为;%v\n" , &a,a)
	a =90
}
func changeInPtr( a *int){
	fmt.Printf("---------cahngeIntVal函数内存\n")
	fmt.Printf("-----a的内存地址为：%p， 值为;%v\n" , &a,a)
	*a = 50
}
/*
值传递：对数据修改不会修改原来的数据
传递参数的地址会变，但值不会变，相当于拷贝了一份
传指针：对数据会改变原来的数据
传参是引用型是可以修改原来的数据（指针,slice,map,chan属于引用类型）
传非引用类型不可以修改原来的数据 （int string bool 数组 strut）
*/