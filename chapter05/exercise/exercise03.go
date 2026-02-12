package main

import "fmt"

func prefixer(prefix string) func(string) string {
	return func(s string) string {
		return prefix + " " + s
	}
}

func exercise03() {
	helloPrefix := prefixer("Hello")
	fmt.Println(helloPrefix("Bob"))   // Hello Bob
	fmt.Println(helloPrefix("Maria")) // Hello Maria
}
