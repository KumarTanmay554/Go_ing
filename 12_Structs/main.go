package main

import "fmt"

type Typename struct {
	Name string
	Age  int
}

func main() {
	var p Typename
	p.Name = "John"
	p.Age = 30
	fmt.Println(p)
}