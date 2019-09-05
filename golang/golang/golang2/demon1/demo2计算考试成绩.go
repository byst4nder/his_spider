package main

import "fmt"

// 函数实参的传递
func main (){
	// 传N个成绩的用法
	sum, avg ,count,user := GetScore( "baven",90,78,85,46,55)
	fmt.Printf("学员%s共有%d门成绩， 总成绩为：%.2f， 平均成绩为%.2f" ,user,count,sum,avg)
	// 传切片的用法
	fmt.Println()
	scores :=[] float64{98, 78, 85, 100,94,86}
	sum, avg, count,user = GetScore( "Tom",scores...)
	fmt.Printf("学员%s共有%d门成绩， 总成绩为：%.2f， 平均成绩为%.2f" ,user,count,sum,avg)
}
// 累加求总分以及品均值    可变参数要放在后面
func GetScore(name string,score ... float64) (sum ,avg float64, count int,user string) {
	fmt.Println("接收到的score为：", score, )
	for _,value := range score {
		sum +=value
		count++
	}
	user =name
	avg = sum /float64(count)
	return
}