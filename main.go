package main

import (
	"fmt"
)

func area(length float64, width float64) float64 {
	return length * width

}

func main() {

	result := area(5.5, 10)
	fmt.Println(result)

}
