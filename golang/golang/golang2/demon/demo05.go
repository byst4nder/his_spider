package main

import "fmt"

type Man struct {
	age, height int
}
func main(){
	const NAME  string = "STEVEN"
	//nam :=Man{30, 177}  不能作为常量
	//const  PERSON  = nam
	const(
		X ="a"
		Y
		A =10
		B
		C
	)
	fmt.Println(A, B, C, X, Y)
	const(
		L =iota
		M =iota
		N =iota
	)
	fmt.Println(L, M, N)
	const(
		L1 =iota
		M1
		N1
	)
	fmt.Println(L1, M1, N1)
	const(
		S = 3
		L2 ="STEVEN"
		M2 =iota
		N2
		F1 = 123
	)
	fmt.Println(L2, M2, N2, S, F1)
	fmt.Println(L, M, N)
	const(
		i = 1 <<iota  // 1*2^   iota
		j = 3 <<iota  // 3*2^1   位运算符
		k
		l
	)
	fmt.Println(i, j, k, l)
	const name  = iota
	fmt.Println(name)
}
