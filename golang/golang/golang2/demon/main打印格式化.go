package main

import "fmt"
 var m init= 19000000
func main() {

	var (
		a int
		b string
		c float32
		d []int
		e int32 =100
		f [3]int
		g bool
		h func() string
		l  =200.34
	)
	m :="19"
	n :="19"
	m, _, q := "a", "b", 10
	// _ 匿名变量，匿名变量不占用空间，
	fmt.Printf("%T , %v\n", a, a)
	fmt.Printf("%T , %q\n", b, b)
	fmt.Printf("%T , %V\n", c, c)
	fmt.Printf("%T , %V\n", d, d)
	fmt.Printf("%T , %V\n", e, e)
	fmt.Printf("%T , %V\n", f, f)
	fmt.Printf("%T , %V\n", g, g)
	fmt.Printf("%T , %V\n", h, h)
	fmt.Printf("%T , %V\n", l, l)
	fmt.Printf("%T , %V\n", m, m)
	fmt.Printf("%T , %V\n", n, n)
	fmt.Printf("%T , %V\n", q, q)

	x :=10
	y :=20
	z :=30
	fmt.Println(x," ", y, " ",z)
	x, y, z = z, x, y
	fmt.Println(x," ", y, " ",z)
	var abc uint16 = 255
	fmt.Println(abc)

}
