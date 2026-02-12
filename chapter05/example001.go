package main

import "fmt"

// 関数の宣言と呼び出し
// func 関数名(引数1 型1, ...) 戻り値の型 { ... } で宣言する
//   - 引数がない場合は単に 「()」 と書く
//   - 同じ引数の型が連続する場合は最後以外の変数の型を省略できる
//     > func div(num int, denom int) { ... }
//     > func div(num, denom int) { ... }
//   - 戻り値を返さない場合は「戻り値の型」に何も書かない
func div(num, denom int) int {
	if denom == 0 {
		return 0
	}
	return num / denom
}

func example001() {
	// 関数の呼び出しは他の言語と共通である
	{
		result := div(5, 2)
		fmt.Println(result)
	}
}
