package main

import (
	"fmt"
	"strconv"
)

// Go の関数は第一級である
//   - 関数を表す型はシグネチャと呼ばれ，キーワード「func」，引数の型，戻り値の型で構成される
//     - 引数と戻り値の型が同じシグネチャは一致しているといえる
//     - type で関数型を定義することもできる
//   - 関数も値なので変数として宣言できる
//   - 関数型のゼロ値は nil であり，nil を実行しようとするとパニックになる
//   - 名前の無い関数 (無名関数・匿名関数) を定義し，編集へ代入することができる
//     - 無名関数は変数を経由せず，そのまま実行できる

func f1(a string) int {
	return len(a)
}

func f2(a string) int {
	total := 0
	for _, v := range a {
		total += int(v)
	}
	return total
}

func _add(i int, j int) int { return i + j }
func _sub(i int, j int) int { return i - j }
func _mul(i int, j int) int { return i * j }
func _div(i int, j int) int { return i / j }

// ↓ 無名関数を用いて，パッケージレベルの変数としてまとめて定義することもできる
var (
	__add = func(i int, j int) int { return i + j }
	__sub = func(i int, j int) int { return i - j }
	__mul = func(i int, j int) int { return i * j }
	__div = func(i int, j int) int { return i / j }
)

func example006() {
	var myFuncVariable func(string) int
	myFuncVariable = f1
	result := myFuncVariable("Hello")
	fmt.Println(result) // 5

	myFuncVariable = f2
	result = myFuncVariable("Hello")
	fmt.Println(result) // 500 // = 72 + 101 + 108 + 111 (ASCIIコードの合計)

	type opFuncType func(int, int) int
	var opMap = map[string]opFuncType{
		"+": _add,
		"-": _sub,
		"*": _mul,
		"/": _div,
	}
	expressions := [][]string{
		[]string{"2", "+", "3"},
		[]string{"2", "-", "3"},
		[]string{"2", "*", "3"},
		[]string{"2", "/", "3"},
		[]string{"2", "%", "3"},
		[]string{"two", "+", "three"},
		[]string{"2", "+", "three"},
		[]string{"5"},
	}
	for _, expression := range expressions {
		if len(expression) != 3 {
			fmt.Print(expression, " -- 不正な式です\n")
			continue
		}
		p1, err := strconv.Atoi(expression[0])
		if err != nil {
			fmt.Print(expression, " -- ", err, "\n")
			continue
		}
		op := expression[1]
		opFunc, ok := opMap[op]
		if !ok {
			fmt.Print(expression, " -- ", "定義されていない演算子です: ", op, "\n")
			continue
		}
		p2, err := strconv.Atoi(expression[2])
		if err != nil {
			fmt.Print(expression, " -- ", err, "\n")
			continue
		}
		result := opFunc(p1, p2)
		fmt.Print(expression, " → ", result, "\n")
	}

	f := func(j int) {
		fmt.Println("無名関数の中で", j, "を出力")
	}
	for i := 0; i < 5; i++ {
		f(i)
	}

	for i := 0; i < 5; i++ {
		func(j int) {
			fmt.Println("無名関数の中で", j, "を出力")
		}(i)
	}

	x := __add(2, 3)
	fmt.Println(x) // 5
	__add = func(i int, j int) int { return i + j + j }
	y := __add(2, 3)
	fmt.Println(y) // 8
}
