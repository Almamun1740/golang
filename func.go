package main 

import "fmt"

func add(x int, y int) int{
	r := x+y

	return r
}
func main(){
	x:=add(10, 30)
	fmt.Println(x)
}