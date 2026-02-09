package main

import "fmt"

func exercise04() {
	var total int
	for i := 0; i < 10; i++ {
		// total := total + i とすると for 文のブロック内で外側の total がシャドーイングされる
		// 以下のようにする (か，total += i とする) と正しく総和を求めることができる
		total = total + i
		fmt.Printf("i=%v total=%v\n", i, total)
	}
	fmt.Printf("total=%v\n", total)
}
