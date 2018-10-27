package main

import (
	"fmt"
)

func main() {
	coursesInProg := []string{"JS", "PHP", "GO"}
	coursesCompleted := []string{"JS", "PHP"}

	for _, i := range coursesInProg {
		for _, j := range coursesCompleted {
			if i == j {
				fmt.Println(i, "completed")
			}
		}
	}

}
