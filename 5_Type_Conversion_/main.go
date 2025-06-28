package main

import (
	"fmt" 
	"strconv"
)

func main(){
	var intnum int8 = 42
	var intnum2 int16 = int16(intnum)
	fmt.Println(intnum2) 

	// var message string = "Hello"
	// var aNum int = int(message) ---not allowed

	var bignum int16 = 300
	var smallnum int8 = int8(bignum)
	fmt.Println(smallnum)
	var afloat float32 = float32(bignum)
	fmt.Println(afloat)

	var message string = "42"
	nums,err := strconv.Atoi(message)
	fmt.Println(nums,err)

	var mess string = strconv.Itoa(nums)
	fmt.Println(mess)
	
}