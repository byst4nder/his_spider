package main

import "fmt"

func main() {
	const (
		a = iota
		b =3
		m =iota
		i
		c
	)
	fmt.Printf("b = %d, c = %d\n",b, c)
	const d  = 3
	const (
		j = iota
		e, f, g  = iota, iota, iota
		l = iota)
	fmt.Printf("e=%d, f=%d, g=%d\n",e, f, g)
	var h bool
	fmt.Println(h)
}
