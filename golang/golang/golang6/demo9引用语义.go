package main

import "fmt"

type Studnet2 struct {
	name string
	age int
}

func (temp *Studnet2)setpoint(a string,b int)  {
	temp.name = a
	temp.age = b
	fmt.Printf("temp = %p\n",temp)
}

func main() {
	s1 :=Studnet2{"张杰",29}
	fmt.Printf("&s1= %p\n",&s1)

	(&s1).setpoint("刘邦",33)
	fmt.Println("s1=",s1)
}
