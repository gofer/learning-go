package main

import "fmt"

func exercise01() {
	greetings := []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}
	greetings1 := greetings[:2]
	greetings2 := greetings[1:4]
	greetings3 := greetings[3:]
	fmt.Println(greetings)
	fmt.Println(greetings1)
	fmt.Println(greetings2)
	fmt.Println(greetings3)
}
