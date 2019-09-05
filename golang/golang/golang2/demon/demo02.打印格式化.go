package main

import "fmt"

type Student struct {
	x,y int
}


/*
[一般]
　　%v	相应值的默认格式。在打印结构体时，“加号”标记（%+v）会添加字段名
　　%#v	相应值的 Go 语法表示
　　%T	相应值的类型的 Go 语法表示
　　%%	字面上的百分号，并非值的占位符
[布尔]
　　%t	单词 true 或 false。
[整数]
　　%b	二进制表示
　　%c	相应 Unicode 码点所表示的字符
　　%d	十进制表示
　　%o	八进制表示
　　%q	单引号围绕的字符字面值，由 Go 语法安全地转义
　　%x	十六进制表示，字母形式为小写 a-f
　　%X	十六进制表示，字母形式为大写 A-F
　　%U	Unicode 格式：U+1234，等同于 "U+%04X"
[浮点数及其复合构成]
　　%b	无小数部分的，指数为二的幂的科学计数法，与 strconv.FormatFloat 的 'b' 转换格式一致。例如 -123456p-78
　　%e	科学计数法，例如 -1234.456e+78
　　%E	科学计数法，例如 -1234.456E+78
　　%f	有小数点而无指数，例如 123.456
　　%g	根据情况选择 %e 或 %f 以产生更紧凑的（无末尾的 0）输出
　　%G	根据情况选择 %E 或 %f 以产生更紧凑的（无末尾的 0）输出
[字符串与字节切片]
　　%s	字符串或切片的无解译字节
　　%q	双引号围绕的字符串，由 Go 语法安全地转义
　　%x	十六进制，小写字母，每字节两个字符
　　%X	十六进制，大写字母，每字节两个字符
[指针]
　　%p	十六进制表示，前缀 0x
[注意]
　　这里没有 'u' 标记。若整数为无符号类型，他们就会被打印成无符号的。类似地， 这里也不需要指定操作数的大小（int8，int64）。
*/
func main()  {
	//通用格式
	str :="steven"
	fmt.Printf("%T, %v\n" , str, str)
	var a rune ='一'
	fmt.Printf("%T, %v\n", a ,a)
	p:=Student{1,2}
	fmt.Printf("%T, %v\n", p, p)
	//布尔格式
	fmt.Printf("%T,  %t\n", true, true)
	//整形格式
	fmt.Printf("%T, %d\n", 123, 123)
	fmt.Printf("%T, %04d\n", 123, 123123)
	fmt.Printf("%T, %.5d\n", 123, 12324523)
	fmt.Printf("%T, %b\n", 123, 123)
	str =fmt.Sprintf("%U\n", '一')
	fmt.Println(str)
	arr :=[3]byte{97,98,99}
	arr1 :=[3]byte{'a','b','c'}
	fmt.Printf("%.2F\n", 123.1124314)
	fmt.Printf("%0.1E\n", 123.1124314)
	fmt.Printf("%s\n", "欢迎大家学习区块连")
	fmt.Printf("%q\n", "欢迎大家学习区块连")
	fmt.Printf("%T , %x \n" , arr, arr)
	fmt.Printf("%T , %X \n" , arr1, arr1)
}



