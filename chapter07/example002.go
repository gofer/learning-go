package main

import (
	"fmt"
	"time"
)

type Counter struct { // Counter型を定義
	total       int       // 更新回数
	lastUpdated time.Time // 最終更新時刻
}

// 型Counterに付随するメソッドIncrement (1増やす) を定義
func (c *Counter) Increment() { // ポインタレシーバー (cはポインタ)
	c.total++
	c.lastUpdated = time.Now()
}

// 型Counterに付随するメソッドStringを定義
func (c Counter) String() string { // 値レシーバー (cにはコピーが渡される)
	return fmt.Sprintf("更新回数: %d, 更新時刻: %v", c.total, c.lastUpdated)
}

func doUpdateWrong(c Counter) { //間違い
	c.Increment() // mainのcのコピーに対してIncrementが行われる
	fmt.Println("Wrong:", c.String())
}

func doUpdateRight(c *Counter) { //正しい
	c.Increment() // mainのcに対してIncrementが行われる
	fmt.Println("Right:", c.String())
}

func example002() {
	var c Counter
	fmt.Println(c.String()) // 更新回数: 0, 更新時刻: 0001-01-01 00:00:00 +0000 UTC
	c.Increment()           // 「(&c).Increment()」と書かなくてもよい
	fmt.Println(c.String()) // 更新回数: 1, 更新時刻: 2026-02-15 22:29:41.702622 +0900 JST m=+0.000290543

	// ポインタのインスタンスのメソッドセットにはポインタレシーバーも値レシーバーも含まれる
	// 値のインスタンスの場合はメソッドセットには値レシーバーのメソッドのみが含まれる
	var c2 Counter
	fmt.Println("main1:", c2.String())
	doUpdateWrong(c2)
	fmt.Println("main2:", c2.String())
	doUpdateRight(&c2)
	fmt.Println("main3:", c2.String())
	doUpdateRight(&c2)
	fmt.Println("main4:", c2.String())
}
