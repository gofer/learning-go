package main

import "fmt"

// 型Personを定義
type Person struct {
	FirstName string // 名
	LastName  string // 姓
	Age       int    // 年齢
}

// 型Personに付随するメソッドStringを定義 (PersonにメソッドStringを付加)
// メソッドは「レシーバー」という特別な引数をもつ関数
// レシーバーpにはPerson型のインスタンスが指定されて呼び出される
func (p Person) String() string { // 「(p Person)」がレシーバーの指定
	return fmt.Sprintf("%s %s：年齢%d歳", p.FirstName, p.LastName, p.Age)
}

func example001() {
	p := Person{ // 上で定義したPerson型の変数pの宣言と初期化
		LastName:  "武田",
		FirstName: "信玄",
		Age:       52,
	}
	output := p.String() // 型Personに付随するメソッドStringを呼び出す
	fmt.Println(output)  // 武田 信玄：年齢52歳
}
