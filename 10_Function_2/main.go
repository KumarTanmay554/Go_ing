package main

import "fmt"

func main() {
	// multipleArgs("Hello", "World", "from", "Go")
	num := 10
	var numPtr *int
	numPtr = &num
	println("Value of address:", numPtr)
	println("Value of address:", &num)

	*numPtr = 20
	println("Value of num:", num)

	a := 10
	b := 20

	pointerFunc(&a, &b)
	fmt.Println("After pointerFunc, a:", a, "b:", b)

}

func multipleArgs(a ...string) {
	for _, v := range a {
		println(v)
	}
}

func pointerFunc(a, b *int) {
	for *a < *b {
		fmt.Println(*a)
		(*a)++
	}
}