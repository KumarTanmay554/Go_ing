package main

import (
	"errors"
	"fmt"
)

func main() {
	printP()
	printP1("Player1")
	ans, err := calc(-1)
	if err!=nil{
		panic(err)
	}
	fmt.Println("Sum of digits:", ans)
}

func printP1(player string){
	fmt.Printf("hello %s \n", player)
}

func printP() {
	fmt.Println("Hello, World!")
}

func calc(a int) (int, error){
	if(a<0){
		return 0, errors.New("negative number not allowed")
	}
	sum:=0
	for a>0{
		sum += a%10
		a/=10
	}

	return sum,nil
}
