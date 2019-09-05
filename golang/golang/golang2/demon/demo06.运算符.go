package main

import "fmt"

var a = 21.0
var b = 21.0 // 定义全局变量
var c float64
func main(){
	//Arithmetic()
	//Relational()
	Logic2()
	Arithmetic()
	Relational()
	Logicl()
}

// 算术运算符
func Arithmetic()  {
	c = a + b
	fmt.Printf("第一行 - c 的值为%.3f\n",c)
	a++
	fmt.Printf("第一行 - a 的值为%.2f\n",int(a))
	b--
	fmt.Printf("第一行 - b 的值为%.2f\n",b)
}

// 关系运算符
func Relational(){
	if (a == b){
		fmt.Printf("a等于b\n")
	}else {
		fmt.Printf("a不等于b\n")
	}
	if (a <= b){
		fmt.Printf("a小于等于b\n")
	}else {
		fmt.Printf("a不小于b\n")
	}
	if (a > b){
		fmt.Printf("a大于b", a, b, "\n")
	}else {
		fmt.Printf("a不大于b\n")
	}
}

// 逻辑运算符 &&  || ！
func Logicl(){
	 a := true
	 b :=false
	if (a && b){
		fmt.Printf("第一行条件成立\n")
	}
	if (a || b){
		fmt.Printf("第二行条件成立\n")
	}
	if (!(a && b)){
		fmt.Printf("第三条条件成立\n")
	}
}

func Logic2(){
	score := 88
	if score >= 60 {
		if score >= 70{
			if score >= 80{
				fmt.Printf("优秀\n")
			}else {
				fmt.Printf("良好\n")
			}
		}else {
			fmt.Printf("及格\n")
		}
	} else {
		fmt.Printf("不及格\n")
	}
}
// 位运算符 Bitwise
// switch 分支语句




