package main

import "fmt"

type Adder struct { // 構造体Adderの定義
	start int // int型のフィールドstartをもつ
}

func (a Adder) AddTo(val int) int { // 型Adderをレシーバーとするメソッドを定義
	return a.start + val // フィールドstartの値に，引数valの値を足して戻す
}

func example004() {
	// インスタンスにメソッドを呼び出している
	myAdder10 := Adder{start: 10}   // startの値を10にして型Adderのインスタンスを生成
	fmt.Println(myAdder10.AddTo(5)) // 15

	// インスタンスのメソッドを「メソッド値」として呼び出している (クロージャーに似ている)
	f1 := myAdder10.AddTo // 型Adderの変数myAdder10のメソッドAddToをf1に代入
	fmt.Println(f1(15))   // 25

	f2 := Adder.AddTo              // 型Adderをレシーバーとして定義されているメソッドAddToをf2に代入
	fmt.Println(f2(myAdder10, 15)) // 25
	// ↑ レスーバーとしてmyAdder10を指定。構文が変わる点に注意!
	// f2のシグネチャは func(Adder, int) int である
}
