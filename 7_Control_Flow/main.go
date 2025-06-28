package main

import (
	"bufio"
	"fmt"
	"os"
	// "strconv"
	"strings"
)

func main(){
	r:= bufio.NewReader((os.Stdin))
	line,err := r.ReadString('\n')
	if err != nil{
		panic(err)
	}

	message := strings.TrimSpace(line)

	// num, err := strconv.Atoi(strings.TrimSpace(line))
	// if err!= nil{
	// 	panic(err)
	// }
	// fmt.Println(num)

	// if num < 10{
	// 	fmt.Println("male")
	// }
	// if num < 20{
	// 	fmt.Println("Female")
	// } else {
	// 	fmt.Println("ohter")
	// }

	switch message{
	case "Monday":
		fmt.Println("Today is Monday")
		fallthrough
	case "Tuesday":
		fmt.Println("Today is Tuesday")
		fallthrough
	case "Wednesday":
		fmt.Println("Today is Wednesday")
		fallthrough
	case "Thursday":
		fmt.Println("Today is Thursday")
		fallthrough
	case "Friday":
		fmt.Println("Today is Friday")
		fallthrough
	case "Saturday":
		fmt.Println("Today is Saturday")
		fallthrough
	case "Sunday":
		fmt.Println("Today is Sunday")
	default:
		fmt.Println("Invalid day")
	}

	for i:=0; i<=10; i++{
		fmt.Println(i)
	}
}