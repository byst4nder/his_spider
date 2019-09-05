package main

import "fmt"

func main()  {
	ScoreGrade()
}
func ScoreGrade(){
	score :=99.5

	switch {
	case score >= 90:
		fmt.Println("优秀")
	case score >=80:
		fmt.Println("良好")
	case score >=70:
		fmt.Println("中等")
	case score >= 60:
		fmt.Println("及格")
	default:
		fmt.Println("差")


	}
}

