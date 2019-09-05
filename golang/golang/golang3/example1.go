package main

import (
	"fmt"
	"time"
)

const (
	Female = 2
	  )

func main()  {
	for{
		second := time.Now().Unix()
		if (second % Female == 0){
			fmt.Println("Female")
		}else{
			fmt.Println("Man")
		}
		time.Sleep(1000 * time.Millisecond)
	}
}
