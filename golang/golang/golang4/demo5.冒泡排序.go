package main

import "fmt"

func main(){
	//arr := []int{4,1,5,3,2}
	arr :=[]int{12,34,2,5,7,6,89,56,99}
	BubbleSort(arr)
	fmt.Println(arr)
}
func BubbleSort(arr []int)  {
	iCount :=0
	jCount :=0
	// 双层for循环实现排序
	for i:=0;i<len(arr)-1;i++{
		flag :=true //定义一个标记，判断是否有两两换位，如果
		for j:=0;j<len(arr)-1-i;j++{
			if arr[j]>arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				// 如果出现两两换位，那么说明排序还没有结束
				flag = false
			}
			jCount++
		}
		iCount++
		if flag{
			break
		}
	}
fmt.Println("i循环的次数：， j循环的次数：" ,iCount,jCount)
}