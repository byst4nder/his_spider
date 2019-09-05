package main

import "fmt"

//for 循环
func main(){
	// 循环方式一
	for i:=0;i<=10;i++{
		fmt.Printf("%d  ",i)
	}
	fmt.Printf("\n")

	// 循环方式二
	j:=0
	for ;; j++{
		if (j > 10) {
			break
		}
		fmt.Printf("%d  ", j)
	}
	fmt.Printf("\n ")

	// 循环方式三
	for ;j <= 20;j++{
		fmt.Printf("%d  ", j)

	}
	fmt.Printf("\n")

	// 循环方式四
	for ;;{
		if j >30{
			break
		}
		j++
		fmt.Printf("%d  ", j)
	}

	// 累加100
	sum :=0
	k :=0
	for ;k<= 100;k++{
		if k % 3 ==0{
			sum +=k;
		}
	}
	fmt.Printf("%d\n", sum)
	//
	count :=0
	long :=32.0
	for i:=0;long > 4;i++{
		long -=1.5
		count +=1
	}
	fmt.Printf("%d", count)
	//
	count1 :=0
	for i:=32.0; i>4; i-= 1.5{
		count1 ++
	}
	// 字符串
	fmt.Printf("%d\n", count1)
	str :="123asdfg阿富汗"
	for i, value:= range str{
		fmt.Printf("第%d的字符是%v  %c\n", i, value, value)
	}
	// 循环数组
	arr :=[]int{123,321,324}
	for i,values := range arr{
		fmt.Printf("%d  %d\n", i, values)
	}
	//
	for i:=0; i<=5;i++{
		for j:=0; j<=5-i ;j++{
			fmt.Printf("*")
		}
		fmt.Println()
	}

}