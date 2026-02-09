package main

import (
	"math/rand"
)

func exercise01() {
	slice := make([]int, 0, 100)
	for i := 0; i < cap(slice); i++ {
		slice = append(slice, rand.Intn(100))
	}
}
