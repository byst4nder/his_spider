package main

import "fmt"

func main()  {
	//创建切片
	numbers :=[]int{0,1,2,3,4,5,5,6,7,8}
	printSlice(numbers)
	number1 := numbers[1:4]
	printSlice(number1)
	Slicecap()
}


func printSlice(x []int)  {
	fmt.Printf("len=%d cap=%d slice=%v \n", len(x), cap(x),x)
}


func Slicecap(){
	arr :=[...]string{"a","b","c","d","e","f","g","h","i","j","k"}
	s1 :=arr[2:8]
	fmt.Printf("%T \n",s1)
	fmt.Println("cap(s1)=", cap(s1),s1)
}
// 切片是引用类型