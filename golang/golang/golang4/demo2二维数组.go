package main

import "fmt"

/*
多维数组
二维数组
*/
func main()  {
	var a = [5][3]int{{1,2,3},{4,5,6},{7,8,9},{10,11,12},{13,14,15}}
	for i:=0;i <len(a);i++{
		for j:=0;j<len(a[0]);j++{
			fmt.Printf("a[%d][%d]=%d\n",i , j, a[i][j])
		}
	}
}




