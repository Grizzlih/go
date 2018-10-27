package main 

import (
	"fmt"
)

func main() {
	testMap := map[string]int{"A": 1, "B": 2}

	testMap["A"] = 100
	testMap["F"] = 77
	delete(testMap, "F")
	fmt.Println(testMap)
}
