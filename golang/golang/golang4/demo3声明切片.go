package main

import "fmt"

// 声明切片
var s1[]int
var s2 []int = make ([]int, 3)
var s3 []int = make([]int,5,7)

func main (){
	s4 := make([]int, 5)
	s5 := make([]int, 5,7)
	s6 := make([]int,0)
	printSliceMsg(s1)
	printSliceMsg(s2)
	printSliceMsg(s3)
	printSliceMsg(s4)
	printSliceMsg(s5)
	if s1==nil{
		fmt.Println("s1 ==nil", "len(s1)", len(s1))
	}
	if s6==nil{
		fmt.Println("s6 ==nil", "len(s6)", len(s6))
	}
	//fmt.Printf("len=%d slice=%d lice=%v \n", len(s2), cap(s2), s2)
}
func printSliceMsg (x []int){
	fmt.Printf("len=%d slice=%d lice=%v \n", len(x), cap(x), x)
}