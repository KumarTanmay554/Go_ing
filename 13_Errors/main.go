package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int,error) {
	if b == 0 {
		return 0, errors.New("division by zero is not allowed")
	}
	return a / b,nil
}

func divideWithPanic(a, b int) int {
	if b == 0 {
		panic("division by zero is not allowed")
	}
	return a / b
}

func main() {
	defer func(){
		if resp:= recover(); resp != nil{
			fmt.Println("Recovered from panic:", resp)
			fmt.Println("Helps server to continue running even after an error.")
		}
	}();

	fmt.Println(divide(10,0))
	fmt.Println("even if error is found, the program continues to run.")

	fmt.Println(divideWithPanic(10,0))
	fmt.Println("This line will not be executed due to the panic above.")

}