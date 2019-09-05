package main

import (
	"fmt"
	"strings"
)

//函数也是一个变量 ，通过type来自定义一个类型，函数完全相同

/*
函数声明包括函数名、形式参数列表、返回值列表、函数体。

*/

func main() {
	str := "abcdesABCDES"
	//result := process(str)
	//fmt.Println(result)
	//函数变量的普通用法
	fmt.Print(StringTocase(str, process))
	fmt.Println()
	// 使用type的函数变量的用法
	fmt.Print(StringTocase2(str, process))
	max_test()
}

// 处理函数实现大小写
func process(str string) string {
	result := ""
	for i, value := range str {
		if i%2 == 0 {
			result += strings.ToUpper(string(value))

		} else {
			result += strings.ToLower(string(value))
		}
	}
	return result
}

func StringTocase(str string, myfunc func(string) string) string {
	return myfunc(str)
}

// 使用type 自定义类型
type caseFunc func(string) string

func StringTocase2(str string, myfunc caseFunc) string {
	return myfunc(str)
}

// 返回两个当中最大的一个
func max(num1, num2 int) int {
	var result int
	if (num1 > num2) {
		result = num1
	} else {
		result = num2
	}
	return result
}

func max_test() {
	a := 100
	b := 200
	ret := max(a, b)
	fmt.Printf("\n最大值为：%d\n", ret)

	d, e := change(3, 4)
	fmt.Println(d, e)
	fmt.Println(sun(1, 2, 3, 4, 5, 6, 7, 8, 9))
/*	values := [...]int{2, 3, 5, 6, 7, 8, 4}
	fmt.Println(sum(values...))*/
}

func change(a, b int) (int, int) {
	return b, a
}

// 可变参数
func sun(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}
