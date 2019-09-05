package main

import (
	"fmt"
	"math"
	"strconv"
)
/*
基本数据类型：整形，浮点型，布尔型，复数，常量 ，字符串
复合数据类型：数组，结构体，slice，map
*/
/*
整型
种类
有符号
int8、int16、int32、int64
无符号
uint8、uint16、uint32、uint64
如果装的系统是32位，则是int32；如果是64则是int64，系统决定使用多少位来存放
Unicode字符rune类型等价int32、byte等价uint8
许多整数之间的相互转换并不会改变数值；它们只是告诉编译器如何解释这个值。但是对于将一个大尺寸的整数类型转为一个小尺寸的整数类型，或者是将一个浮点数转为整数，可能会改变数值或丢失精度
当使用fmt包打印一个数值时，我们可以用%d、%o或%x参数控制输出的进制格式
*/

func test() {
	var u uint8 = 255
	fmt.Println(u, u+2, u*u)
	var i int8 = 127
	fmt.Println(i, i+3, i*i)

	var a int32 = 1
	var b int16 = 2
	var comp = int(a) + int(b)
	fmt.Println(comp)

	ascii := 'a'
	unicode := '国'
	newline := '\n'
	fmt.Printf("%d %[1]c %[1]q\n", ascii)   // "97 a 'a'"
	fmt.Printf("%d %[1]c %[1]q\n", unicode) // "22269 国 '国'"
	fmt.Printf("%d %[1]q\n", newline)       // "10 '\n'"
	var f float32 = 16777216                // 1 << 24
	fmt.Println(f == f+1)                   // "true"!
	for x := 0; x < 8; x++ {
		fmt.Printf("x = %d e^x = %8.3f\n", x, math.Exp(float64(x)))
	}
}

func main() {
	a := 100
	var b byte = 100
	var c rune = 200
	var e byte = 'a'
	var f rune = '一'
	var str string = "abcde"
	/*
	   T 打印出字符类型  v表示格式化
	*/
	fmt.Printf("%T, %v\n", a, a)
	fmt.Printf("%T, %v\n", b, b)
	fmt.Printf("%T, %v\n", c, c)
	fmt.Printf("%T, %v\n", e, e)
	fmt.Printf("%T, %v\n", f, f)
	fmt.Printf("%T, %v\n", str, str)

	temp := `
    x :=10
	y :=20
	z :=30
	fmt.Println(x," ", y, " ",z)
	x, y, z = z, x, y
	fmt.Println(x," ", y, " ",z)
	var abc uint16 = 255
	fmt.Println(abc)
`
	fmt.Println(temp)
	test()
	complex_test()
	string_test()
	const_test()
	arr_test()
	slince_test()
	map_test()
	struct_test()
}

/*
浮点数
Go语言提供了两种精度的浮点数，float32和float64。
复数
Go语言提供了两种精度的复数类型：complex64和complex128，分别对应float32和float64两种浮点数精度。内置的complex函数用于构建复数，内建的real和imag函数分别返回复数的实部和虚部
布尔型
一个布尔类型的值只有两种：true和false。
布尔值并不会隐式转换为数字值0或1，反之亦然。必须使用一个显式的if语句辅助转换：
字符串
常量
一个常量的声明语句定义了常量的名字，和变量的声明语法类似，常量的值不可修改，这样可以防止在运行期被意外或恶意的修改
*/

/*
符合数据类型
Go语言中有四种复合数据类型：数组，slice，map，结构体
Go语言中有四种复合数据类型：数组，slice，map，结构体
数组和结构体是聚合类型；它们的值由许多元素或成员字段的值组成。数组是由同构的元素组成——每个数组元素都是完全相同的类型——结构体则是由异构的元素组成的。数组和结构体都是有固定内存大小的数据结构。相比之下，slice和map则是动态的数据结构，它们将根据需要动态增长。
数组

*/

func complex_test() {
	var x complex128 = complex(1, 2)
	var y complex128 = complex(3, 4)
	fmt.Println(x * y)
	fmt.Println(real(x * y))
	fmt.Println(imag(x * y))
}

// 将整数转化成为字符串
func string_test() {
	x := 123
	y := fmt.Sprintf("%d", x)
	fmt.Println(y, strconv.Itoa(x), x)
	// 用16进制来格式化数字，FormatInt和FormatUint函数可以用不同的进制来格式化数字
	fmt.Println(strconv.FormatInt(int64(x), 16))

}

func const_test() {
	const (
		a = iota
		b
		c
		d = "haha"
		e
		f = 100
		g
		h = iota
		i
		j = 8 << iota
	)
	fmt.Println(a, b, c, d, e, f, g, h, i, j)

}

/*
复合数据类型
Go语言中有四种复合数据类型：数组，slice，map，结构体
数组和结构体是聚合类型，它们的值由许多元素或成员字段的值组成
数组是由同构的元素组成——每个数组元素都是完全相同的类型
结构体则是由异构的元素组成的
数组和结构体都是有固定内存大小的数据结构。相比之下，slice和map则是动态的数据结构，它们将根据需要动态增长
*/

/*
数组
数组是一个由固定长度的特定类型元素组成的序列，一个数组可以由零个或多个元素组成。因为数组的长度是固定的，因此在Go语言中很少直接使用数组。和数组对应的类型是Slice（切片），它是可以增长和收缩动态序列，slice功能也更灵活，但是要理解slice工作原理的话需要先理解数组。
数组的每个元素可以通过索引下标来访问，索引下标的范围是从0开始到数组长度减1的位置。内置的len函数将返回数组中元素的个数。
*/

func arr_test() {
	fmt.Println("数组练习部分")
	var a [3]int
	a = [3]int{1, 6}
	fmt.Println(a[0])
	fmt.Println("数组的长度为：", len(a)-1)
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}
	var q = [3]int{1, 2, 4}
	var r = [3]int{1, 2}
	b := [3]int{1, 4, 5}
	c := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(q, r, b, c)

	e := [...]int{99: -1}
	fmt.Println(e)

	fmt.Println("数组的相互比较")
	p := [2]int{1, 2}
	l := [...]int{1, 2}
	v := [2]int{1, 3}
	//m := [...]int{1, 2}
	//d := [3]int{1, 2}
	fmt.Println(p == l, l == v, p == v)
	//fmt.Println(m == d)
}

/*
Slice（切片）代表变长的序列，序列中每个元素都有相同的类型。一个slice类型一般写作[]T，其中T代表slice中元素的类型；slice的语法和数组很像，只是没有固定长度而已。
一个slice由三个部分构成：指针、长度和容量。指针指向第一个slice元素对应的底层数组元素的地址，要注意的是slice的第一个元素并不一定就是数组的第一个元素。长度对应slice中元素的数目；长度不能超过容量，容量一般是从slice的开始位置到底层数据的结尾位置。内置的len和cap函数分别返回slice的长度和容量。
len()切片的长度， cap()切片的容量
*/

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// 判断连个数组是否相等
func equal(x, y []string) bool  {
	if len(x) != len(y){
		return false
	}
	for i:= range x{
		if x[i] != y[i]{
			return false
		}
	}
	return true
}

func slince_test() {
	months := [...]string{1: "January", 12: "December"}
	Q2 := months[4:7]
	summer := months[6:9]
	fmt.Println(Q2)     // ["April" "May" "June"]
	fmt.Println(summer) // ["June" "July" "August"]
	a := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	reverse(a[:])
	fmt.Println(a)
	var runes []rune
	for _, r :=range "helle, 世界" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes)
}
/*
内置的make函数创建一个指定元素类型、长度和容量的slice。容量部分可以省略，在这种情况下，容量将等于长度。
*/

func selice_copy() {
	var numbers []int
	printSlice(numbers)

	/* 允许追加空切片 */
	numbers = append(numbers, 0)
	printSlice(numbers)

	/* 向切片添加一个元素 */
	numbers = append(numbers, 1)
	printSlice(numbers)

	/* 同时添加多个元素 */
	numbers = append(numbers, 2,3,4)
	printSlice(numbers)

	/* 创建切片 numbers1 是之前切片的两倍容量*/
	numbers1 := make([]int, len(numbers), (cap(numbers))*2)

	/* 拷贝 numbers 的内容到 numbers1 */
	copy(numbers1,numbers)
	printSlice(numbers1)
}

func printSlice(x []int){
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}

/*
Map
Map 是一种无序的键值对的集合。Map 最重要的一点是通过 key 来快速检索数据，key 类似于索引，指向数据的值。
可以使用内建函数 make 也可以使用 map 关键字来定义 Map
*/

func map_test()  {
	ages := map[string]int{
		"alice": 31,
		"charlie": 34,
	}
	ages["age"] = 50
	fmt.Println(ages["alice"])
	delete(ages, "charlie") // 删除某个值

	// 便利map中的key/value
	for name, age := range ages{
		fmt.Printf("%s\t%d\n", name, age)
	}
	age, ok :=ages["age"]
	if !ok{
		fmt.Println("存在当前元素", age, ok)
	}else {
		fmt.Println("存在", age, ok)
	}
}

/*
结构体
Go 语言中数组可以存储同一类型的数据，但在结构体中我们可以为不同项定义不同的数据类型。
结构体是由一系列具有相同类型或不同类型的数据构成的数据集合。
如果结构体的全部成员都是可以比较的，那么结构体也是可以比较的，那样的话两个结构体将可以使用==或!=运算符进行比较。相等比较运算符==将比较两个结构体的每个成员，因此下面两个比较的表达式是等价的
*/

type Books struct {
	title string
	auther string
	subject string
	book_id int
}
type Point struct {
	X, Y int
}

type address struct {
	hostname string
	port     int
}

func struct_test()  {
	fmt.Println("结构体练习")
	fmt.Println(Books{"Golang 语言", "huang", "Go语言教程", 6495407})
	fmt.Println(Books{title:"GO语言", auther:"huang"})
	p :=Point{1, 2}
	q :=Point{2, 1}
	fmt.Println(p.X ==q.X, p.X == q.Y)
	fmt.Println(p == q)

	hits := make(map[address]int)
	hits[address{"golang.org", 443}]++
}