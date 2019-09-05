package main

import "fmt"

func main(){
	// 传n个值
	fmt.Println(AddSun(1,2,3,4,5,6))
	// 不传值
	fmt.Println(AddSun())
	// 传切片
	arr :=[]int{3,4,5,34,53,36}
	fmt.Println(AddSun(arr ...)) // 穿切片 后面追加三个... 表示对其进行解包
}
// 累加求和，参数是可变参数，参数个数从0—n
func AddSun(nums ...int) ( sum int){
	fmt.Printf("%T\n" , nums)
	for _,value:= range nums{
		sum += value
	}
	return
}
