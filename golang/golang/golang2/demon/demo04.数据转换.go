package main

import "fmt"

func main(){ //  数据类型的转换  注意数据的精度  布尔型无法和其他类型进行转换
	chinses := 90
	english := 80.9
	avg := (float64(chinses) + english)/2
	fmt.Println(avg)
	//flag := true
	//string(flag)
	//str := "1313"
	//int(str)
	x := 1233
	str := string(x) // 转化成了
	fmt.Println(str)
	fmt.Printf("%T, %v\n", str, str)
}

// 常量的定义 const 标示符 [类型] = 值
// const WIDTH,HEIGHT =  valuel, value2
/*
常量组  如果不指定类型和初始值， 则与上一行非空常量的值相同。
const
iota  特殊常量值   是一个系统定义的常量值可以被修改 iota只能出现在常量组中。
iota可以理解成常量组中的常量的计数器，不论该常量的值是多少，只要有一个常量，那么iota就加1
*/

