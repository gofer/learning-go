package main

import "fmt"

// Go は値渡しの言語である
//   - つまり，関数に引数を渡すとコピーが作られる
//   - スライスやマップは実装としてポインタが用いられているため，引数として渡された変数が変更される
//   - ただし，スライスのサイズを変更することはできない

type person struct {
	age  int
	name string
}

func modifyFails(i int, s string, p person) {
	i *= 2 // 「i = i * 2」と同じ
	s = "さようなら"
	p.name = "Bob"
}

func modMap(m map[int]string) {
	m[2] = "こんにちは"
	m[3] = "さようなら"
	delete(m, 1)
}

func modSlice(s []int) {
	for k, v := range s {
		s[k] = v * 2
	}
	s = append(s, 10)
}

func example009() {
	p := person{}
	i := 2
	s := "こんにちは"
	fmt.Println(i, s, p) // 2 こんにちは {0 }
	modifyFails(i, s, p)
	fmt.Println(i, s, p) // 2 こんにちは {0 }

	m := map[int]string{
		1: "1番目",
		2: "2番目",
	}
	modMap(m)
	fmt.Println(m) // map[2:こんにちは 3:さようなら]

	t := []int{1, 2, 3}
	modSlice(t)
	fmt.Println(t) // [2 4 6]
}
