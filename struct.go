package main

import "fmt"

type Book struct{
	Title string
	Author string
	ISBN string
	Price float32
	Pages int
}

func main(){
	var b1 Book
	b1.Title="An Introduction to programming"
	b1.Author="CALEB DOXY"
	b1.ISBN="978-1478355823"
	b1.Price=10.50
	b1.Pages=165

	fmt.Println(b1)
	fmt.Println(b1.Title)
	fmt.Println(b1.Author)
	fmt.Println(b1.ISBN)
	fmt.Println(b1.Price)
	fmt.Println(b1.Pages)
}