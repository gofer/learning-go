package main

import "fmt"

type Doubler interface {
	Double()
}

type DoubleInt int

// ポインタレシーバでintの値を変更する
func (di *DoubleInt) Double() {
	*di = *di * 2
}

// スライスは比較可能ではないことに注意する
type DoubleIntSlice []int

// スライス型では値レシーバが使える
func (dis DoubleIntSlice) Double() {
	for i := range dis {
		dis[i] = dis[i] * 2
	}
}

// Doubler型の2つの引数を受け取り，等しいかどうかを出力する
func DoublerCompare(d1, d2 Doubler) {
	fmt.Println(d1 == d2)
}

func example009() {
	var di DoubleInt = 10
	var di2 DoubleInt = 10
	var dis = DoubleIntSlice{1, 2, 3}
	// var dis2 = DoubleIntSlice{1, 2, 3}

	// ↓ これが false になるのは，値ではなくポインタを比較しており，異なるインスタンスであるから
	DoublerCompare(&di, &di2) // false

	// ↓ これが false になるのは，型が異なるから
	DoublerCompare(&di, dis) // false

	// ↓ コンパイルはされるが，実行時にパニックになる
	// DoublerCompare(dis, dis2) // panic: runtime error: comparing uncomparable type main.DoubleIntSlice
}
