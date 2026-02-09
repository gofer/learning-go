/**
 * 4章 ブロック、シャドーイング、制御構造
 */

package main

import (
	"fmt"
	"math/rand"
	"unicode/utf8"
)

func main() {
	// ブロック: 変数や関数などの宣言が行われる場所をブロックと呼ぶ
	//   - 関数の外で宣言された変数，定数，型，関数はパッケージレベルに置かれる
	//   - import文によって他のファイルからインポートされた識別子はファイルブロックに置かれる
	//   - 関数のトップレベルで定義されているすべての変数 (関数の引数を含む) はひとつのブロックに含まれる
	//   - 関数の中では { ... } によって別のブロックを定義できる
	//   - Goの制御構造 (for, if など) も独自のブロックを定義する
	//   - 内側のブロックは外側のブロックで定義された任意の識別しにアクセスできる
	// シャドーイング: 内側のブロックで外側のブロックと同じ名前の識別子を定義すると，外側の識別子が隠ぺいされてしまう
	{
		x := 10
		if x > 5 {
			fmt.Println(x) // 10 (外側のx)
			x := 5
			fmt.Println(x) // 5 (内側のx，外側のxがシャドーイングされた)
		}
		fmt.Println(x) // 10
	}
	// := で変数宣言をすると簡単に変数をシャドーイングできてしまう
	// 以下は複数変数への代入によるシャドーイングの例
	{
		x := 10
		if x > 5 {
			a, x := 5, 20
			fmt.Println(a, x) // 5 20 (内側のa, x，外側のxがシャドーイングされた)
		}
		fmt.Println(x) // 10 (外側のx)
	}
	// パッケージ名さえもシャドーイングできてしまう
	{
		x := 10
		fmt.Println(x)
		fmt := "おっと〜"
		// ↓ 実行時エラーになってしまう
		// fmt.Println(x) // fmt.Println undefined (type string has no field or method Println)
		// 以下では fmt は string型の変数であり，fmtパッケージはシャドーイングされている
		println(fmt)
	}
	// ユニバースブロック: Goはキーワードが25個しかなく，事前宣言された型 (int, stringなど) や定数 (true, false など) ，
	//   組み込み関数はユニバーサルブロックに事前宣言された識別子として定義されている
	//   - よってこれらもシャドーイングされる可能性がある
	//   - ユニバースブロックにある識別子は再定義すべきではない
	{
		fmt.Println(true) // true
		true := 10
		fmt.Println(true) // 10 (trueがシャドーイングされた)
	}

	// if: 条件分岐を行う制御構造
	// if <条件式> { ... } else if <条件式> { ... } else { ... } の形
	{
		n := rand.Intn(10) // 0以上10未満の整数を戻す
		if n == 0 {
			fmt.Println("少し小さすぎます:", n)
		} else if n > 5 {
			fmt.Println("大きすぎます:", n)
		} else {
			fmt.Println("いい感じの数字です:", n)
		}
	}
	// Goのif文では，条件式および各if/elseブロックの中で有効な変数を宣言できる
	//   - というより，任意の単純な文を書くことができる (が，混乱の原因なので非推奨である)
	{
		if n := rand.Intn(10); n == 0 {
			fmt.Println("少し小さすぎます:", n)
		} else if n > 5 {
			fmt.Println("大きすぎます:", n)
		} else {
			fmt.Println("いい感じの数字です:", n)
		}
		// ↓ はコンパイルできない (nのスコープから外れている)
		// fmt.Println(n) // undefined: n
	}

	// for: 繰り返しを行う制御構造
	//   - Goには繰り返しを行う制御構造はforしかない
	//   - Goには4種類のfor文がある
	// 1. 標準形式のfor文
	//   for <初期化>; <条件>; <再設定> { ... }
	//     - 初期化では複数の変数を初期化することもできる
	//     - 条件は評価結果がbool型になる式を指定する
	//     - 条件式が評価されるタイミングは，初期化後のループが実行される前と，再設定が実行された後である
	//     - 再設定は任意の代入ができ，各ループの後，条件が評価される前に実行される
	{
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
	}
	// 3つのパート (初期化，条件，再設定) はいずれかも省略できる
	{
		i := 0
		for ; i < 10; i++ {
			fmt.Println(i)
		}
		for j := 0; j < 10; {
			fmt.Println(j)
			if j%2 == 0 { // 2で割ってあまりが0
				j++ // jを1増やす
			} else {
				j += 2 // jを2増やす
			}
		}
	}
	// 2. 条件式のみのfor文 (while 相当)
	// 条件式のみを記述すれば他言語の while 相当になる
	// 初期化と再設定の両方を省略する場合は ; を書かない
	{
		i := 1
		for i < 100 {
			fmt.Println(i)
			i = i * 2
		}
	}
	// 3. 無限ループのfor文
	// 条件式も初期化も再設定も省略すれば無限ループになる
	{
		// 無限ループするのでコメントアウト
		// for {
		// 	 fmt.Println("Hello")
		// }
	}
	// break と do-while 相当のfor文
	// for 文の中でループを抜け出すには break を使う
	// また，do-while 相当のループを実現するには条件式のみのfor文と break を組み合わせる
	//     do { <処理> } while (<条件>);
	//   は以下のように書ける
	//     for { <処理>; if !<条件> { break } }
	//   (!<条件> はループの継続条件の否定が break の条件になるためである)
	// for 文の中でループの最中に次のループを開始するためには continue を使う (条件は再評価される)
	{
		// continue を使わずに書いたコード
		// このようなコードはイディオム的ではない (Goでは if の本文は短くし，ネストを浅くすることが推奨されている)
		for i := 1; i <= 100; i++ {
			if i%3 == 0 {
				if i%5 == 0 {
					fmt.Println(i, "3でも5でも割り切れる")
				} else {
					fmt.Println(i, "3で割り切れる")
				}
			} else if i%5 == 0 {
				fmt.Println(i, "5で割り切れる")
			} else {
				fmt.Println(i)
			}
		}

		// continue を使ってネストを浅くしたコード
		for i := 1; i <= 100; i++ {
			if i%3 == 0 && i%5 == 0 {
				fmt.Println(i, "3でも5でも割り切れる")
				continue
			}
			if i%3 == 0 {
				fmt.Println(i, "3で割り切れる")
				continue
			}
			if i%5 == 0 {
				fmt.Println(i, "5で割り切れる")
				continue
			}
			fmt.Println(i)
		}
	}
	// 4. for-range ループ
	// 事前宣言されている型の各要素に対して繰り返しを行うための構文
	//   for <インデックス変数>, <要素変数> := range <コレクション> { ... }
	//   - 配列，スライス，文字列，マップ，チャネルに対して使える
	//   - インデックス変数は配列やスライス，文字列だと i (index) が，マップだと k (key) がよく使われる
	//   - 要素変数は v (value) がよく使われる
	{
		evenVals := []int{2, 4, 6, 8, 10, 12} // 偶数
		for i, v := range evenVals {
			fmt.Println(i, v)
			// 0, 2
			// 1, 4
			// 2, 6
			// ...
		}
	}
	// インデックス変数や要素変数を使わない場合は _ (アンダースコア) を使う
	{
		evenVals := []int{2, 4, 6, 8, 10, 12} // 偶数
		for _, v := range evenVals {
			fmt.Println(v)
			// 2
			// 4
			// 6
			// ...
		}
	}
	// キーは必要はだが，値はいらない場合は2つ目の変数を省略できる
	{
		uniqueNames := map[string]bool{"花子": true, "太郎": true, "洋子": true}
		for k := range uniqueNames {
			fmt.Println(k)
			// 花子
			// 太郎
			// 洋子
		}
	}
	// マップのイテレーションでは順序は保証されない (Hash DoS対策でハッシュ関数に乱数が使われている)
	// (ただし，fmt.Println などではキーでソートされる)
	{
		m := map[string]int{"a": 1, "c": 3, "b": 2}
		for i := 0; i < 3; i++ {
			fmt.Println("ループ", i)
			for k, v := range m {
				fmt.Println(k, v)
			}
		}
	}
	// 文字列のイテレーションではバイト列ではなく rune に対してイテレーションを行う
	//   - もし，UTF-8として認識できないバイトに出会うと 0xFFFD が返る
	{
		samples := []string{"hello", "apple_π!", "これは漢字文字列"}
		for _, sample := range samples {
			for i, r := range sample {
				fmt.Println(i, r, string(r))
			}
			fmt.Println()
		}
	}
	// for-range ループの値はコピーである
	//   - Go 1.22より前では，値のコピーは一度だけ生成されて，各ループで同じ変数が再利用されていた
	//   - Go 1.22以降では，各ループで新しい変数が生成されるようになった
	{
		evenVals := []int{2, 4, 6, 8, 10, 12}
		for _, v := range evenVals {
			v *= 2
		}
		fmt.Println(evenVals) // [2 4 6 8 10 12]
	}
	// break や continue は既定では一番内側のループに対して適用される
	// ラベルを利用することで，一気に多重ループを抜け出せる
	{
		samples := []string{"hello", "apple_π!", "これは漢字文字列"}
	outer: // outer ラベル
		for _, sample := range samples {
			for i, r := range sample {
				fmt.Println(i, r, string(r))
				if r == 'l' || r == 'は' {
					continue outer
				}
			}
			fmt.Println()
		}
	}
	// 各 for 文の選択の指針
	//   - 最も利用されるのは for-range ループ
	//     - スライスやマップ，チャネルで利用される
	//     - 特に文字列はバイトではなくrune単位
	//   - 3つ指定の標準 for 文は全要素に対してイテレーションしない場合に便利
	//     - ただし，標準形式の for 文では文字列はバイト列として扱われるので，「最初のn文字を飛ばす〜」などは正しく扱えない
	//   - それ以外のループの使用頻度は小さい
	{
		// 以下の2つはどちらも「最初から2番目」から「最後から2番目」までを処理する
		evenVals := []int{2, 4, 6, 8, 10}

		// for-range 版
		for i, v := range evenVals {
			if i == 0 {
				continue
			}
			fmt.Println(i, v)
			if i == len(evenVals)-2 {
				break
			}
		}

		// 標準形式の for 版
		for i := 1; i < len(evenVals)-1; i++ {
			fmt.Println(i, evenVals[i])
		}
	}

	// switch文: 多分岐の条件分岐を行う制御構造
	// Goでは他言語とは異なる強力な構文がある
	// 1. 単純な switch 文
	//   - switch の式には switch 文全体で使える変数を宣言できる
	//   - case 節には複数の値をカンマ区切りで列挙できる
	//   - case 節 (default 節を含む) の中で定義した変数はその case 節内でのみ有効
	//   - case 節は既定でフォールスルーしない (非推奨だが明示的に fallthrough を書くとフォールスルーする)
	{
		words := []string{"山", "sun", "微笑み", "人類学者", "モグラの穴", "mountain", "タコの足とイカの足", "antholopologist", "タコの足は8本でイカの足は10本"}
		for _, word := range words {
			switch rc := utf8.RuneCountInString(word); rc { // Rune count
			case 1, 2, 3, 4:
				fmt.Printf("「%s」の文字数は%dで，短い単語だ。\n", word, rc)
			case 5:
				bc := len(word) // Byte count
				fmt.Printf("「%s」の文字数は%dで，これはちょうどよい長さだ。ちなみにバイト数は%dだ。\n", word, rc, bc)
			case 6, 7, 8, 9: // "mountain", "タコの足とイカの足" など 6〜9文字の単語では何も出力しない
			default:
				fmt.Printf("「%s」の文字数は%dで，とても長い！。\n", word, rc)
			}
		}
	}
	// ブランク switch: switch 文の比較対象の変数などを指定しないもの
	// ブランク switch では case 節に任意の論理式を指定する
	{
		words := []string{"hi", "salutations", "hello"}
		for _, word := range words {
			switch wordLen := len(word); {
			case wordLen < 5:
				fmt.Println(word, "は短い単語です")
			case wordLen > 10:
				fmt.Println(word, "は長すぎる単語です")
			default:
				fmt.Println(word, "はちょうどよい長さの単語です")
			}
		}
	}
	// switch 文内の break
	//   - switch 文の中で break が必要な場合はほとんどない
	//   - switch 文の中で外側のループを抜け出すためにラベルへ break することはある
	{
		// ラベルを指定しない break
		for i := 0; i < 10; i++ {
			switch {
			case i%2 == 0:
				fmt.Println(i, ":偶数")
			case i%3 == 0:
				fmt.Println(i, ":3で割り切れるが2では割り切れない")
			case i%7 == 0:
				fmt.Println(i, ":ループ終了…はしない！")
			default:
				fmt.Println(i, ":退屈な数")
			}
		}
		// 7 :ループ終了…はしない！
		// 8 :偶数
		// 9 :3で割り切れるが2では割り切れない

		// ラベルを指定した break
	loop:
		for i := 0; i < 10; i++ {
			switch {
			case i%2 == 0:
				fmt.Println(i, ":偶数")
			case i%3 == 0:
				fmt.Println(i, ":3で割り切れるが2では割り切れない")
			case i%7 == 0:
				fmt.Println(i, ":ループ終了！")
				break loop
			default:
				fmt.Println(i, ":退屈な数")
			}
		}
		// 7 :ループ終了！
	}
	// if 文と switch 文 の使い分け: 各ケースで比較する値に何らかの関連がある場合は switch 文のほうが可読性が上がることもある
	{
		// if 文を利用した場合
		for i := 1; i <= 100; i++ {
			if i%3 == 0 && i%5 == 0 {
				fmt.Println(i, "3でも5でも割り切れる")
				continue
			}
			if i%3 == 0 {
				fmt.Println(i, "3で割り切れる")
				continue
			}
			if i%5 == 0 {
				fmt.Println(i, "5で割り切れる")
				continue
			}
			fmt.Println(i)
		}

		// switch 文を利用した場合
		for i := 1; i <= 100; i++ {
			switch {
			case i%3 == 0 && i%5 == 0:
				fmt.Println(i, "3でも5でも割り切れる")
			case i%3 == 0:
				fmt.Println(i, "3で割り切れる")
			case i%5 == 0:
				fmt.Println(i, "5で割り切れる")
			default:
				fmt.Println(i)
			}
		}
	}
	// goto 文: 指定したラベルへジャンプできる
	//   - goto 文は可読性を著しく損なうため利用は可能な限り避けるべきである
	{
		a := rand.Intn(10)
		for a < 100 {
			fmt.Println(a)
			if a%5 == 0 {
				goto done
			}
			a = a*2 + 1
		}
		fmt.Println("ループが通常終了したときに行う処理を実行")
	done:
		fmt.Println("ループが終わったときに必ず行う処理を実行")
		fmt.Println(a)
	}
	//   - Goの goto 文はジャンプできる場所に制限がある
	//     * 変数の宣言を超えるジャンプはできない
	//     * 内側のブロックの中や並列しているブロックの中にはジャンプできない
	{
		// 不正な goto 文
		a := 10
		// goto skip　// goto skip jumps over declaration of b at ./main.go:405:5
		b := 20
		// skip:
		c := 30
		fmt.Println(a, b, c)
		if c > 1 {
			// goto inner // goto inner jumps into block starting at ./main.go:412:12
		}
		if a < b {
			// inner:
			fmt.Println("aはbより小さい")
		}
	}
}
