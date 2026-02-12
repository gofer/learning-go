package main

import (
	"errors"
	"fmt"
	"os"
)

// Go の関数は複数の戻り値を返せる
//   - その場合は戻り値の型を (, ) で囲み，その中に型をカンマ区切りで列挙する
//   - return での戻り値は括弧で括らないことに注意する
//   - 関数内で起きたエラーは戻り値の最後で戻し，正常終了の場合は nil を返すのが慣例
//   - ただし「タプル」のような概念ではないので，戻り値を受け取る際は一つずつ変数に受け取る必要がある
//   - 戻り値を受け取る変数が不要な場合は，ブランク識別子 _ を使って無視できる
//   - 全ての戻り値を無視する場合は，そもそも受け取らなければ良い

func divAndRemainder(num, denom int) (int, int, error) {
	if denom == 0 {
		return 0, 0, errors.New("0で割ることはできません")
	}
	return num / denom, num % denom, nil
}

func example004() {
	result1, remainder, err := divAndRemainder(5, 2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(result1, remainder) // 2 1

	result2, _, _ := divAndRemainder(5, 2)
	fmt.Println(result2, remainder)
}
