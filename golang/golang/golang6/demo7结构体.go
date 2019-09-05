package main

import (
	"fmt"

)

type Person struct {
	class int
	id int
}

type Student struct{
	Person
	name string
	age int
	adr string
	mer bool
}

func test(t *Student) {
	(*t).age = 23
	fmt.Println("t=",*t )
}

/*
顺序初始化和指定成员初始化
*/

func main() {
	// 顺序初始化
	p1 := Student{Person{1,33},"张三",12,"深圳",true}
	s := &p1
	test(s)
	// fmt.Println("p1=",p1)
	fmt.Println("p1",p1)
	fmt.Printf("p1 = %+v\n",p1)
	// 指定成员初始化
	p2 :=Student{Person:Person{class:1},name :"李四"}
	fmt.Println("p2=",p2)
}