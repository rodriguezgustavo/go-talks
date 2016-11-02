package main

import (
	"fmt"
)

func main() {
	arrayType := [5]string{"1", "2", "3", "4", "5"}

	for _, value := range arrayType {
		fmt.Print(value+" ")
	}
}
