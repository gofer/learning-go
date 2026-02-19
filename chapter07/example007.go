package main

type LogicProvider struct{}

func (lp LogicProvider) Process(data string) string {
	// ビジネスロジック
	return data
}

type Logic interface {
	Process(data string) string
}

type Client struct {
	L Logic
}

func (c Client) Program() {
	data := "input data" // どこからかデータを取得
	c.L.Process(data)
}

func example007() {
	c := Client{
		L: LogicProvider{},
	}
	c.Program()
}
