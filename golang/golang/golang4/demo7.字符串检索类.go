package main

import (
	"fmt"
	"strings"
)

func main(){
	a :=strings.Contains("1.sesfood","foo")
	fmt.Println(a)
	fmt.Println(strings.ContainsAny("2.abcd","ia"))
	fmt.Println(strings.ContainsRune("3.tean",'定'))
	fmt.Println(strings.Count("4.fiadf","f"))
	fmt.Println(strings.HasPrefix("5.10000gd","100"))
	fmt.Println(strings.HasSuffix("6.fwe123","123"))
	fmt.Println(strings.Index("7.abcdef","cd"))
	fmt.Println(strings.IndexAny("8.abcdef","坑的几e"))
	fmt.Println(strings.IndexByte("9.123bac",'a'))
	fmt.Println(strings.IndexRune("10.大荣教育",'教'))
	// fmt.Println(strings.IndexFunc("helloafhj中", func(r rune) bool {}))
}
