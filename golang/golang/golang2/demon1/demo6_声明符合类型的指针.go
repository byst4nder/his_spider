package main

import "fmt"

type Student struct {
	name string
	age int
	married bool
	sex int8
}

func main(){
	s1 :=Student{name :"Steven", age :35, married:true, sex:1}
	s2 :=Student{name:"Sunny", age:20, married:false, sex:0}
	var a *Student = &s1
	var b *Student = &s2
	fmt.Println("\n-----------")
	fmt.Printf("s1类型%T, 值为%v \n" ,s1, s1)
	fmt.Printf("s2类型%T, 值为%v \n" ,s2, s2)
	fmt.Printf("------------------\n")
	fmt.Printf("a类型%T, 值为%v \n" ,a, a)
	fmt.Printf("*b类型%T, 值为%v \n" ,*b, *b)
	fmt.Println(&a.name, &a.age, &a.married, &a.sex, &a)
}

//空指针 当没有给一个指针分配任何变量时，他的值为nil
// 一个指针变量通常缩写为ptr
// if (ptr ！= nil) ptr 不是空指针
// if (ptr == nil) ptr 是空指针










