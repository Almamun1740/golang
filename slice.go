package main

import "fmt"
import "reflect"

func main() {
	//var students [3]string
	//students[0]="Mamun"
	//students[1]=""
	//students[2]=""
	//x := students[0:3]

	//x := make([]string, 0)
	var fruits []string
	fruits = append(fruits, "Apple", "Banana", "Mango")
	//fmt.Println(fruits, len(fruits))

	//fmt.Printf("%T\n", fruits)
	//fmt.Printf("%T\n", students)
	
	b := reflect.TypeOf(fruits).Kind().String()
	fmt.Println(b)
}