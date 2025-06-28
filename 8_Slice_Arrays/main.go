package main

import "fmt"

func main() {
	slicess := []int{1, 2, 3, 4, 5}
	// fmt.Println(slicess)
	slicess[1] = 10
	fmt.Println(slicess)
	// fmt.Println(slicess[0])

	// for i:=0; i<len((slicess)); i++{
	// 	slicess[i] += 1
	// }
	// fmt.Println(slicess)

	// for i,element := range slicess{
	// 	fmt.Println("index",i,"element",element)
	// 	slicess[i] += 1
	// }
	// fmt.Println(slicess)
	

	arr := [5]int{1,2,3,4,5}
	fmt.Println(arr)
	// no append in arrays
	// fmt.Println(arr[0])
	// arr[1] = 10
	// fmt.Println(arr) 

	// MAPS
	mp := map[string]string{
		"key1":"1",
		"key2":"2",
		"key3":"3",
		"key4":"4",
	}
	// fmt.Println(mp)
	// fmt.Println(mp["key1"])
	// mp["key1"] = "10"
	// fmt.Println(mp)

	// hp, ok:= mp["key5"]
	// fmt.Println(hp, ok) // ok will be false if key5 does not exist

	// mp["key5"] = "5" 
	// fmt.Println(mp)

	// delete(mp, "key5")
	// fmt.Println(mp)

	for key, value := range mp{
		fmt.Println("key:", key, "value:", value)
	}
	// random order in maps
}