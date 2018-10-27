package main

import (
	"fmt"
)

func main() {
	myCourses := []string{"JS", "PHP"}
	fmt.Printf("Length is: %d. \nCapacity is: %d", len(myCourses), cap(myCourses))

}
