package main

import "fmt"

type Student4 struct {
	name string
	age int
}

func (temp Student4)setinfo()  {
	fmt.Println("Student4:",temp.name, temp.age)
}

type Student5 struct {
	Student4
	height int
}
// 匿名字段方法的值可以重写
func (temp Student5)setinfo()  {
	fmt.Println("Student5:",temp.name, temp.height)
}

func main() {
	s3 :=Student5{Student4{"小强",18},180}
	s3.setinfo()
	// 显示调用继承方法
	s3.Student4.setinfo()
	// 方法值
	pfnc :=s3.setinfo
	pfnc()
	//方法表达式
	pfnc11 :=(Student5).setinfo
	pfnc11(s3)
}
