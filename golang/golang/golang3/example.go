package main

import "fmt"
// for 循环练习
func list(n int)  {
	for i := 0;i <=n; i++{
		fmt.Printf("%d+%d=%d\n", i, n-i, n)
	}
}

func main(){
	list(10)
}
