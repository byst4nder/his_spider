package main

import "fmt"

// strings  strconv
func main(){
	s :="我爱go语言"
	fmt.Println("字节的长度：", len(s))  // 字符串是字节的长度
	// 遍历所有字符串
	for i,ch :=range s {
		fmt.Printf("%d:%c \n",i, ch)
	}
	// 遍历所有字节
	for i,ch :=range []byte(s){
		fmt.Printf("%d:%x \n",i,ch)
	}
	//遍历所有字符
	for i,ch := range []rune(s){
		fmt.Printf("%d:%x:%v \n",i,ch, ch)
	}
}
/*
字符串是指每个汉字，字节
*/