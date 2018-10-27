package main

import (
	"fmt"
	"os"
)

func main() {

	name := os.Getenv("USER")
	course := "Docker Deep Dive"
	const speedOfLightMph = 186000

	fmt.Println("\nSpeed of light mph", speedOfLightMph)

	fmt.Println("\nHi", name, "you're currently watching", course)

	changeCourse(&course)

	fmt.Println("\nYou are now watching course", course)
}

func changeCourse(course *string) string {

	*course = "First Look: Native Docker Clustering"

	fmt.Println("Trying to change your couse to", *course)

	return *course
}
