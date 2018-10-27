package main

import (
	"fmt"
)

func main() {

	leagueTitles := make(map[string]int)
	leagueTitles["Sunderland"] = 6
	leagueTitles["Newcastle"] = 4

	recentHead2Head := map[string]int{
		"Sunderland": 5,
		"Newcastle":  0,
	}

	fmt.Printf("%v\n%v", leagueTitles, recentHead2Head)

}
