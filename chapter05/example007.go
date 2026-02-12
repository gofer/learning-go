package main

import (
	"fmt"
	"sort"
)

// クロージャー: 関数が定義された環境への参照を保持する仕組み
//   - 関数を引数として関数を渡せる
//     - 渡された関数はそれ自身が定義された環境を参照できるため，局所変数を外に持ち出すことができる
//   - 関数を戻り値として戻せる

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func makeMult(base int) func(int) int {
	return func(factor int) int {
		return base * factor
	}
}

func example007() {
	a := 20
	f := func() {
		fmt.Println(a)
		a = 30
	}
	f()            // 20
	fmt.Println(a) // 30

	people := []Person{
		{"Pat", "Patterson", 37},
		{"Tracy", "Bobbert", 23},
		{"Fred", "Fredson", 18},
	}
	fmt.Println("●初期データ")
	fmt.Println(people)

	// 姓 (LastName) でソート
	sort.Slice(people, func(i, j int) bool {
		return people[i].LastName < people[j].LastName
	})
	fmt.Println("●姓 (LastName。2番目のフィールド) でソート")
	fmt.Println(people)

	// 年齢 (Age) でソート
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println("●年齢 (Age) でソート")
	fmt.Println(people)

	fmt.Println("●ソート後のpeople")
	fmt.Println(people) // sort.Sliceは元のスライスを変更する

	twoBase := makeMult(2)   // 2倍する関数
	threeBase := makeMult(3) // 3倍する関数
	for i := 0; i <= 5; i++ {
		fmt.Print(i, ": ", twoBase(i), ", ", threeBase(i), "\n")
	}
}
