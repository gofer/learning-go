package main

import (
	"fmt"
	"math/rand"
)

func exercise02() {
	slice := make([]int, 0, 100)
	for i := 0; i < cap(slice); i++ {
		slice = append(slice, rand.Intn(100))
	}

	for _, v := range slice {
		switch {
		case v%6 == 0:
			fmt.Println("Six!")
		case v%2 == 0:
			fmt.Println("Two!")
		case v%3 == 0:
			fmt.Println("Three!")
		default:
			fmt.Println("Never mind")
		}
	}
}
