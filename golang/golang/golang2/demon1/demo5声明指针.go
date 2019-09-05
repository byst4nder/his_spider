package main

import "fmt"

func main()  {
	//实际变量
	a :=120
	// 声明指针变量
	var ip *int
	//给指针变量赋值
	ip =&a
	//打印类型
	fmt.Printf("a 的类型是%T，值是%v \n", a, a) //int 120
	fmt.Printf("&a 的类型是%T，值是%v \n", &a, &a) //地址
	fmt.Printf("ip 的类型是%T，值是%v \n", ip, ip) //地址
	fmt.Printf("ip 的类型是%T，值是%v \n", *ip, *ip) //int 指针前面加*，表示原来的指针变量
	fmt.Println(a, &a, *&a,"\n" ,ip, &ip, *(&ip), &(*ip) )

	}
/*
在实际变量前加&表示取变量的指针，在指针前加*表示取指针的变量值，在指针前加&表示取指针的指针
*/
