package main

import "fmt"

// Counter と Increment の再掲 (example002.go)

// type Counter struct { // 構造体Counterの定義
// 	 total       int
// 	 lastUpdated time.Time
// }

// func (c *Counter) Increment() { // *Counterをレシーバーとするメソッドの定義
// 	 c.total++
// 	 c.lastUpdated = time.Now()
// }

// func (c Counter) String() string { // Counterをレシーバーとするメソッドの定義
// 	 return fmt.Sprintf("total: %d, lastUpdated: %v", c.total, c.lastUpdated)
// }

type Incrementer interface {
	Increment() // 関数IncrementをもてばIncrementerになれる
}

func example008() {
	var pointerCounter *Counter        // 構造体Counterを指すポインタ。ゼロ値はnil
	fmt.Println(pointerCounter == nil) // true // 値はnil (ゼロ値) なので
	var incrementer Incrementer        // インタフェースIncrementerを満たす変数の定義。ゼロ値はnil
	fmt.Println(incrementer == nil)    // true
	incrementer = pointerCounter       // pointerCounterは*CounterなのでIncrementの実装をもつので，incrementerに代入できる
	fmt.Println(incrementer == nil)    // false // pointerCounter
}
