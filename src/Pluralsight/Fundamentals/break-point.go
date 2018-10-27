package main

import (
	"fmt"
)

func main() {
	scores := []int{1, 2, 3, 4, 5, 6}
	achievedScores := []int{4, 6}
	completedScores := []int{4, 6}

breakPoint:
	for _, i := range scores {
		fmt.Println("verify")
		for _, j := range achievedScores {
			if i == j {
				fmt.Println("achieved")
				for _, g := range completedScores {
					if j == g {
						fmt.Println("finished")
						break breakPoint
					}
				}
			}
		}
	}
}
