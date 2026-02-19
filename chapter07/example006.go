package main

import (
	"fmt"
)

type Inner struct {
	A int
}

func (i Inner) IntPrinter(val int) string {
	return fmt.Sprintf("Inner: %d", val)
}

func (i Inner) Double() string {
	result := i.A * 2
	return i.IntPrinter(result)
}

type Outer struct {
	Inner // 埋め込まれた構造体
	S     string
}

func (o Outer) IntPrinter(val int) string {
	return fmt.Sprintf("Outer: %d", val)
}

// o.Double() を呼び出すと，InnerのDoubleメソッドが呼び出され，その中のInner.IntPrinterが呼び出される
// レシーバーがOuterであっても，Innerとして処理される
func example006() {
	o := Outer{
		Inner: Inner{
			A: 10,
		},
		S: "Hello",
	}
	fmt.Println(o.Double()) // Inner: 20
}
