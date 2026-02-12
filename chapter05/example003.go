package main

import "fmt"

// 可変長引数とスライス
// 可変長引数を使う場合は，関数の最後の引数の型の前に「...」を付ける
// 可変長引数は関数内では指定された型のスライスとなる
// 可変長引数にスライスを渡すこともできるが，その際はスライスの後ろに「...」を付ける必要がある

func addTo(base int, vals ...int) []int {
	out := make([]int, 0, len(vals))
	for _, v := range vals {
		out = append(out, base+v)
	}
	return out
}

func example003() {
	fmt.Println(addTo(3))             // []
	fmt.Println(addTo(3, 2))          // [5]
	fmt.Println(addTo(3, 2, 4, 6, 8)) // [5 7 9 11]

	a := []int{4, 3}
	fmt.Println(addTo(3, a...))                    // [7 6]
	fmt.Println(addTo(3, []int{1, 2, 3, 4, 5}...)) // [4 5 6 7 8]
	// ↓ はコンパイルできない
	// fmt.Println(addTo(3, a)) // cannot use a (variable of type []int) as int value in argument to addTo
	// ↓ はコンパイルできない
	// fmt.Println(addTo(3, []int{1, 2, 3, 4, 5})) // cannot use []int{…} (value of type []int) as int value in argument to addTo
}
