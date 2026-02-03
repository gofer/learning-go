package main

import "fmt"

func exercise02() {
	const value = 3141592
	var i int = value
	var f float64 = value

	fmt.Println(i, f) // 3141592 3.141592e+06
}
