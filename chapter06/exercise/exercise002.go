package main

import "fmt"

func UpdateSlice(ss []string, s string) {
	ss[len(ss)-1] = s
	fmt.Println("UpdateSlice:", ss)
}

func GrowSlice(ss []string, s string) {
	ss = append(ss, s)
	fmt.Println("GrowSlice:", ss)
}

func exercise002() {
	ss := []string{"a", "b", "c"}
	fmt.Println("before:", ss)
	UpdateSlice(ss, "x")
	// こちらは正しく変更される
	fmt.Println("after update:", ss)
	GrowSlice(ss, "y")
	// こちらは変更されない (長さは変えられない)
	fmt.Println("after grow:", ss)
}
