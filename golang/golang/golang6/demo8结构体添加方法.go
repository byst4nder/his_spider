package main

import "fmt"

type Student1 struct {
	name string
	age int
	height1 int
}
//带有接收者的函数叫做方法
func (temp Student1) Printinfo()  {
	fmt.Println("temp=",temp)
}
//通过一个函数给，给成员赋值
func (p *Student1) Setinfo(a string,b int,c int)  {
	p.name = a
	p.age = b
	p.height1 = c
}


func main() {
	s1 :=Student1{"张三",23,170}
	s1.Printinfo()
	var s2 Student1
	// 修改成员函数
	(&s2).Setinfo("李四",22,179)
	s2.Printinfo()
}