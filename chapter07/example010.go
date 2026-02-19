package main

import (
	"errors"
	"fmt"
)

type treeVal interface {
	isToken()
}

type number int

func (number) isToken() {}

type operator func(int, int) int

func (operator) isToken() {}

func (o operator) process(n1, n2 int) int {
	return o(n1, n2)
}

var operators = map[string]operator{
	"+": func(n1, n2 int) int { return n1 + n2 },
	"-": func(n1, n2 int) int { return n1 - n2 },
	"*": func(n1, n2 int) int { return n1 * n2 },
	"/": func(n1, n2 int) int { return n1 / n2 },
}

type treeNode struct {
	val            treeVal
	lchild, rchild *treeNode
}

func walkTree(t *treeNode) (int, error) {
	switch val := t.val.(type) {
	case nil:
		return 0, errors.New("不正な式")
	case number:
		// t.valがnumber型だとわかったのでint型の値を返す
		return int(val), nil
	case operator:
		// t.valがoperator型だとわかったので左右の子を取得し
		// operatorのメソッドprocess()を呼び出し，値を処理した結果を返す
		left, err := walkTree(t.lchild)
		if err != nil {
			return 0, err
		}
		right, err := walkTree(t.rchild)
		if err != nil {
			return 0, err
		}
		return val.process(left, right), nil
	default:
		// treeValに新しい型が定義されたが，walkTreeは更新されていないことがわかる
		return 0, errors.New("不明な型のノード")
	}
}

func example010() {
	// 1 + 2 * 3
	val, err := walkTree(&treeNode{
		val: operators["+"],
		lchild: &treeNode{
			val: number(1),
		},
		rchild: &treeNode{
			val: operators["*"],
			lchild: &treeNode{
				val: number(2),
			},
			rchild: &treeNode{
				val: number(3),
			},
		},
	})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(val) // 7
}
