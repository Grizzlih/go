package main

import (
	"fmt"
)

func main() {
	firstRank := "39"
	secondRank := "614"
	if firstRank < secondRank {
		fmt.Println("\n First course is doing than Second course")
	} else if firstRank > secondRank {
		fmt.Println("\n Second course is doing than First course")
	} else {
		fmt.Println("\n Both courses the same")
	}
}
