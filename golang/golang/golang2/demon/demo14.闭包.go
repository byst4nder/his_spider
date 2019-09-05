package main
import "fmt"
func main(){
	res := func()(func() int){
		i :=10
		return func() int {
			i++
			return i
		}
	}()
	fmt.Println(res())
}
