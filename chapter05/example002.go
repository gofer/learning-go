package main

import "fmt"

// Go には名前付き引数やオプション引数が存在せず，原則全ての引数を渡す必要がある
// 各引数に対応するフィールドを持った構造体を定義して渡せば同等のことを実現できる
// (とはいえ関数の引数は少ないほうがよく，それらが必要だと思うならば関数自体が複雑すぎるのである)

type MyFuncOpts struct { // MyFunc のオプション引数
	FirstName string // 名
	LastName  string // 姓
	Age       int    // 年齢
}

func MyFunc(opts MyFuncOpts) error {
	fmt.Println(opts)
	fmt.Println("〈ここで必要な処理を行う〉")
	return nil
}

func example002() {
	MyFunc(MyFuncOpts{
		LastName: "Patel",
		Age:      50,
	}) // { Petel  50}
	MyFunc(MyFuncOpts{
		FirstName: "Joe",
		LastName:  "Smith",
	}) // {Joe Smith 0}
}
