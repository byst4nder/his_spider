package main

import (
	"fmt"
)

func main() {
	fmt.Printf("a = %c , a = %d",'a','a')
	fmt.Printf("a = %c , a = %d","a","a")
	fmt.Printf("a = %c , a = %d",'a'-32,'a')
	var b complex64 = 22 + 33i +44 +55
	fmt.Println("实部为：",real(b),"虚部为：",imag(b))
	fmt.Println("请输入num的值：")
    fmt.Scanf("%d",&b)
	fmt.Println(b)
	num := 83
	switch  {
	case num > 90:
		fmt.Println("优秀")
	case num >80:
		fmt.Println("良好")
	case num >70:
		fmt.Println("及格")
	default:
		fmt.Println("不及格")
	}
}
