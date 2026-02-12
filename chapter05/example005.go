package main

import (
	"errors"
	"fmt"
)

// 名前付き戻り値
//   - 戻り値の型の前に識別子を指定することで戻り値に名前を付けられる
//   - 名前を付けた戻り値はゼロ値で初期化される
//   - 戻り値の一部に名前を付けたい場合は _ を名前を付けない戻り値の識別子として利用する
// 名前付き戻り値のデメリット: 可読性が下がるので，利用しないことが推奨される
//   - 名前付き戻り値を利用するとシャドーイングの可能性がある
//   - 名前付き戻り値を利用しても return で別の値を返せる
//   - ブランク return という全く何も値を返さない return が可能になってしまう

func divAndRemainder2(num int, denom int) (result int, remainder int, _ error) {
	result, remainder = 20, 30 // 適当な値を代入
	if denom == 0 {
		return result, remainder, errors.New("0で割ることはできません")
	}
	result, remainder = num/denom, num%denom
	return result, remainder, nil
	// return 10, 20, nil // この return だと result, remainder ではなく，10, 20 が返される
}

func example005() {
	rs, rm, _ := divAndRemainder2(5, 2)
	fmt.Println(rs, rm) // 2 1
}
