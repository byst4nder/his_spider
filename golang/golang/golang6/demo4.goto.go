package main

import "fmt"

func MyFunc1(args ...int)  {
	fmt.Printf("MyFunc len =%d\n ",len(args))
	for i :=0;i<len(args);i++{
		fmt.Printf("MyFunc[%d]=%d\n",i, args[i])
	}
	for i,data := range args  {
		fmt.Printf("Myfunc[%d]=%d\n",i, data)
	}
	fmt.Println(args)
}
func MyFunc3(args ...int)  {
	for i,data :=range args{
		fmt.Printf("Myfunc[%d]=%d\n",i, data)
	}
}
func MyFunc2(args ...int){
	MyFunc3(args[:len(args)] ...) // 前闭后开
}
func MyFunc4(args ...int)(max,min int){
	for i:=0;i<len(args)-1;i++{
		if args[i]<args[i+1]{
			max = args[i+1]
			min = args[i]
		}else {
			max = args[i]
			min = args[i+1]
		}
	}
	return
}
func main() {
	fmt.Println("11111111")
	goto End
	fmt.Println("222222222")
	End:
	fmt.Println("333333333")
	MyFunc1(1,3,5)
	MyFunc2(1,2,3,4,6,7,8,9)
	a,b := MyFunc4(1,4,5,8,9) // 解包
	fmt.Println("max=",a, "min=",b)
}
