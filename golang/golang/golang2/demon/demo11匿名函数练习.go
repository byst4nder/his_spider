package main

import (
	"fmt"
	"math"
)
// 匿名函数作为回调函数的用法
type myFuncs func(float642 float64) string
func main()  {
	// 定义将一个切片，对其中的数据进行求平方跟和求平方的运算
	arr := []float64{1,9,16,25,30}
	result :=FilterSlice(arr , func(val float64) string {
		val = math.Sqrt(val)
		return fmt.Sprintf("%.2f" , val)
	})
	fmt.Print(result)
	// 求平方运算
	//arr := []float64{1,9,16,25,30}
	result =FilterSlice(arr ,func(val float64) string {
		val = math.Pow(val , 2)
		return fmt.Sprintf("%.0f" , val)
	})
	fmt.Print(result)
}
// 遍历切片，对其中每个元素进行运算处理
func FilterSlice(arr [] float64, f myFuncs) [] string{
	var result [] string
	for _ ,value := range arr{
		result = append(result , f(value))
	}
	return result
}
/*
闭包 Closure：闭包不是什么 闭包在形式上像一个函数 lanbda 表达式
函数是执行的行为 ，属性是行为， 而闭包是附有数据的行为。
闭包的价值： 加强模块化，抽象，简化代码
一个编程语言需要哪些特性来支持闭包？
函数可以作为另一个函数的返回值，还可以作为一个变量的值。函数可以嵌套定义
允许定义匿名函数。
*/