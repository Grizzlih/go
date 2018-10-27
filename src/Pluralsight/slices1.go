package main

import (
	"fmt"
)

func main() {
	mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(mySlice[4])

	mySlice[4] = 0
	fmt.Println(mySlice)

	sliceOfSlice := mySlice[2:5]

	fmt.Println(sliceOfSlice)
	sliceOfSlice[1] = 100
	fmt.Println(mySlice)
}
