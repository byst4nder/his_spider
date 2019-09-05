package main

import "fmt"
func test_goroute(a int, b int)  {
	var sum int
	sum = a + b
	fmt.Println(sum)
}

func test_print(a int)  {
	fmt.Println(a)
}

func main()  {
	test_goroute(2,3)
	test_print(5)
}