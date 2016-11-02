package main

import (
	"fmt"
)

func main() {
	mapType := make(map[int]int)

	mapType[1] = 1
	mapType[2] = 2
	mapType[3] = 3
	mapType[4] = 4
	mapType[5] = 5

	for key, value := range mapType {
		fmt.Printf("Key: %d, Value: %d \n", key, value)
	}
}
