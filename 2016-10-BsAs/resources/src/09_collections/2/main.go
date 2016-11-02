package main

import (
	"fmt"
)

func main() {
	sliceType := make([]int, 4, 4)

	sliceType = append(sliceType, 1)
	sliceType = append(sliceType, 2)
	sliceType = append(sliceType, 3)
	sliceType = append(sliceType, 4)
	sliceType = append(sliceType, 5)

	for _, value := range sliceType {
		fmt.Print(value)
	}
}
