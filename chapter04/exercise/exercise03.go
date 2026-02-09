package main

import "fmt"

func exercise03() {
	var total int
	for i := 0; i < 10; i++ {
		total := total + i
		fmt.Printf("i=%v total=%v\n", i, total)
	}
	fmt.Printf("total=%v\n", total)
}
