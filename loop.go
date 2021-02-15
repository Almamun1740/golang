package main 

import (
	"fmt"
)

func main() {
	
	fmt.Print("Enter your age: ")
	var age int
	fmt.Scanf("%d", &age)

	//if boolean_expression {
	//logic or statement here
	//}
	if age < 3 {
		fmt.Println("infant")
	}else if 5>2 && 5<13 {
		fmt.Println("children")
	}else if age>12 && age <= 19{
		fmt.Println("teen age")
	}else{
		fmt.Println("adult")
	}
 	
	//fixed value
	/*
	switch age {
	case 2: 
		fmt.Println("infant")
	case 3,4,5,6,7,8,9,10,11,12:
		fmt.Println("children")
	default: 
		fmt.Println("adult")
	}
	*/

	//fmt.Println(age)
	
	//for loop
	/* for i:1; i<9; i++ {
	fmt.Println(i)
	}
	*/
	//array string literals
	/*
	students := []string{"Al","Mamun","Khan"}
	for i, std := range students {
	fmt.Println(i, std)
	}
	*/
	
	i:=0
	for i<10 {
		fmt.Println(i,"Hello")
	i++
	}

}