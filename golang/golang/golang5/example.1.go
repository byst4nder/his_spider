package main

import "fmt"

func main(){
	for{
		var cmd string
		fmt.Println("请输入数据")
		fmt.Scanf("%s", &cmd)
		cmd1 :=string(cmd)
		fmt.Println(cmd1)
	}
}


