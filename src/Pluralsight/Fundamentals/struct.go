package main

import (
	"fmt"
)

func main() {
	type courseMeta struct {
		Author string
		Level  string
		Rating float64
	}

	//var DockerDeepDive courseMeta
	//DockerDeepDive := new(courseMeta)
	DockerDeepDive := courseMeta{
		Author: "Ed",
		Level:  "Advanced",
		Rating: 5,
	}

	fmt.Println("\n Author is:", DockerDeepDive.Author)
	fmt.Println("\n Rating is:", DockerDeepDive.Rating)

}
