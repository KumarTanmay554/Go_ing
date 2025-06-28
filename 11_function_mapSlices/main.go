package main

import "fmt"

func main() {
	// maps are always pointers in Go
	mp := map[string]int{"foo": 1}
	updateMap(mp)
	fmt.Println("Updated map:", mp)

	sl := []int{1,2,3,4}
	updateSlice(sl)
	fmt.Println("Updated slice:", sl)
}

func updateMap(mp map[string]int) {
	mp["bar"] = 2
}

func updateSlice(sl []int) {
	for i:=range sl{
		sl[i] *=2
	}
}