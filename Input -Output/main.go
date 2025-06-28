package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
) 

func main(){
	// fmt.Println("Hello, World!")
	// fmt.Println("Hello, World!")
	// fmt.Println("Hello, World!")

	// fmt.Println("hello","kjhg")


	// input from
	r:= bufio.NewReader(os.Stdin)
	line ,err := r.ReadString('\n')

	if err != nil{
		panic(err)
	}

	fmt.Println(line)
 
	// num, err := strconv.Atoi(strings.TrimSpace(line))

	// if err != nil{
	// 	panic(err)
	// }
	// fmt.Println("output from atoi",num)

	floatnum , err := strconv.ParseFloat(strings.TrimSpace(line),64)
	if err != nil{
		panic(err)
	}

	fmt.Println("float number",floatnum)
}