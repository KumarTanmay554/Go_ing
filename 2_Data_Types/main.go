package main

import (
	"fmt"
	// "time"
)

func main(){

	// PRIMITIVES DATA TYPES

	// whole number - 14
	// var number int = 14
	// fmt.Println(number)


	// int == int64  *****
	// byte == uint8  ******


	// var num int16 = 14
	// var num3 int8 = 127
	// var num2 int32 = 140000000000 --- too big for int32
	// fmt.Println(num)
	// fmt.Println(num3)

	// var neg uint = -14 -- only positive number with uint
	// var a uint8 = 255
	// fmt.Println(a)


	// real numbers - 1.43
	// var real float64 = 1.43   //15digits after decimal
	// var integer float32 = 1 //7digits after decimal
	// fmt.Println(real)
	// fmt.Println( integer)

	// characters - S, !
	// var character rune = 'S'
	// fmt.Println(character)  // not print "S" but will print "83" as it is mapped with "S"

	//  boolean -true flase
	// var booleanval bool = true
	// fmt.Println(booleanval)

	// text - hello world
	// var text string = "hello world"
	// fmt.Println(text)

	// var result time.Time = time.Now()
	// fmt.Println(result)

	
	
	
	
	num1:=14
	num2:=24
	// sum:=num1+num2
	// sub:=num1-num2
	// mul:=num1*num2
	// div:=num2/num1
	// var divi float32 = float32(num1) / float32(num2)
	// mod:=num2%num1
	
	// fmt.Println(sum)
	// fmt.Println(sub)
	// fmt.Println(mul)
	// fmt.Println(div)
	// fmt.Println(divi)
	// fmt.Println(mod)
	
	num1++
	num2--
	fmt.Println(num1)
	fmt.Println(num2)
	
	// var ch rune = 'A'
	// var ch2 rune = 'B'
	// fmt.Println(ch," ", ch2)
	// fmt.Println(ch+ch2," ", ch-ch2, " ",ch*ch2, " ",ch/ch2, " ",ch%ch2,"Operation on characters")
	
	
	var smallint int8 = 100
	smallint = smallint + 28 //--overflow
	fmt.Println(smallint)
	
	smallint = smallint-2
	fmt.Println(smallint)





	// COMPOSITE DATA TYPES

	var array [5]int = [5]int{1,2,3,4,5}
	fmt.Println(array)
}