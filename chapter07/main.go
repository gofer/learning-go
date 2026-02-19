/**
 * 7章 型、メソッド、インタフェース
 */

package main

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func main() {
	// 抽象型と具象型: 型は全て何をするかを定義するが，どのようにするかも規定するのが具象型，規定しないのが抽象型である
	//   - Goにおける抽象型はインタフェース，具象型は事前宣言された型や構造体・配列・スライス・マップ・チャネルなどそれ以外のものである
	// 基底型: Goのユーザー定義型を含む全ての型は基底型をもっている
	//   - 事前宣言された型はそれ自体が基底型である
	//   - type MyInt int や type YourInt MyInt のように宣言された型は基底型がintである
	// 型定義: Goの型定義は type 識別子 型定義とする
	{
		// 型リテラルを使ってこの構造体リテラルを基底型としてもつ型Personを定義
		type Person struct {
			FirstName string // 名
			LastName  string // 姓
			Age       int    // 年齢
		}

		// intを基底型として型Scoreを定義
		type Score int

		// 型stringを引数，型Scoreを戻り値とする関数型Converterを定義
		type Converter func(string) Score

		// 型stringをを型Scoreにマップするmap型TeamScoresを定義
		type TeamScores map[string]Score
	}
	// メソッド: 定義した型に付随する関数を定義したもの
	//   - func (レシーバー 型) メソッド名(...引数) 戻り値型 { ... } という構文で定義する
	//   - メソッドはパッケージレベルでのみ定義できる (関数は任意のブロックで定義できる)
	//   - レシーバーの名前は型の名前を簡略化したものを使うのが慣習 (this, selfなどはイディオム的でない)
	//   - メソッド名はオーバーロード (多重定義) できない (ある型で複数の同一名メソッドは定義できない)
	example001()
	// ポインタ型レシーバーと値型レシーバー
	//   - メソッドがレシーバーの値を変更する必要がある場合はポインタレシーバーを使う
	//   - レシーバーがnilの可能性がある場合はポインタレシーバーを使う
	//   - それ以外の場合は値レシーバーを利用しても構わない
	//   - Goではゲッターメソッドやセッターメソッドは原則書かず，直接フィールドへアクセスする
	example002()
	// nilインスタンスへの対応
	//   - レシーバーがnilの場合，値レシーバーはパニックになるが，ポインタレシーバーでは直ちにパニックにはならない
	//   - 場合によってはnilレシーバーを許容するほうがコードが単純になる (多くはない)
	//   - メソッドに渡されるレシーバーのポインタはポインタのコピーなのでオリジナルは変更されない
	//     - つまり，nilを受け取ってnilではないものに変更することはできない
	example003()
	// メソッドも関数: メソッドも関数であり，特定の型に関連付けられているということ以外は同様である
	//   - メソッド値: インスタンスのメソッドは値として扱え，関連付けられたインスタンスにアクセスできる
	//   - メソッド式: 型のメソッドから作成した関数で，最初の引数にレシーバーをとる (シグネチャが変わる)
	example004()
	// 関数で実装するか，メソッドで実装するかは他のデータに依存するかで決める
	//   - パッケージレベルの状態はイミュータブルであるべき
	//   - ロジックが入力引数のみに依存するなら関数で実装する
	// 型宣言と継承の違い
	//   - 新しい型を定義する歳には事前宣言された型や構造体リテラルの他に，すでに自分で定義したユーザー定義型をベースにすることもできる
	//   - 多くのオブジェクト指向言語では継承を中心に据えるが，Goではある型を基底型にすることは継承ではない
	//   - Goでは同じ基底型を持っていても，代入などは型変換が必要である (型の間に階層関係はない)
	//   - リテラル・定数・演算子について: 事前宣言された型T1を基底型としてもつユーザー定義型T2に対して
	//     - T1 に代入できるリテラルや定数を T2 に代入できる
	//     - 型 T1 の変数等に関して使える演算子は型 T2 に対しても使える
	{
		type Score int       // intから新たな型(ユーザー定義型) Scoreを定義
		type HighScore Score // ユーザー定義型Scoreから別のユーザー定義型HighScoreを定義

		type Person struct { // ユーザー定義型Person (人) を定義
			LastName  string // 姓
			FirstName string // 名
			Age       int    // 年齢
		}
		type Employee Person // Personを使ってEmployee (従業員) を定義

		// 型のない定数の代入は認められている
		var i int = 300
		var s Score = 100 // Scoreの基底型はintなので100は代入できる
		var hs HighScore = 200
		// hs = s // コンパイル時のエラー！型が違う (ScoreとHighScore)
		// s = i // コンパイル時のエラー！型が違う (Scoreとint)
		s = Score(i)       // 型変換後に代入
		hs = HighScore(s)  // 型変換後に代入
		fmt.Println(s, hs) // 300 300
		hhs := hs + 20     // 基底型 (int) に対して使える演算子 (+) は使える
		fmt.Println(hhs)   // 320
		var s2 Score = 50
		scoreWithBonus := s2 + 10   // scoreWithBonusの型はScore
		fmt.Println(scoreWithBonus) // 60
	}
	// 型は実行可能なドキュメント: 事前宣言された型を基底として新しい型を作るのは，その概念に名前を与えるため
	//   - たとえばint型より，Percentage型のほうが意味が理解しやすい
	//   - あるユーザー定義型を基底として別のユーザー定義型を宣言する場合にも同様のことが言える
	// iotaと列挙型: Goには列挙型がないが，iotaを使って徐々に増加する値を一連の定数に割り与えられる
	//   - iotaは0オリジンである
	//     - 0を利用したくないならconstブロック内で _ = iota で始めることもできる
	//   - iotaの値はconstブロック内の各定数ごとに増加する (使用されているかは関係ない)
	//   - iotaを使うのは明示的に値が定義されておらず，内部用途で利用される場合にのみ用いる
	//   - 値に意味がある場合は仕様書・設計書等に記載されている値を直接書くべきである
	//   - iotaを利用するときはすでに定義されている識別子の後ろに追加する
	//     - 途中に入れると，その後の値がすべて変化してしまうことに注意
	{
		// 型 MailCategory を定義
		type MailCategory int // メールの分類

		// 取りうる値の集合を定義
		const (
			Uncategorized MailCategory = iota // 未分類
			Personal                          // 個人的
			Spam                              // 迷惑メール
			Social                            // ソーシャル
			Advertisement                     // 広告
		)

		const (
			Field1 = 0
			Field2 = 1 + iota // constの2行目なのでiotaは1になる
			Field3 = 20       // 明示的に値が指定されるとその値になる
			Field4            // 値の指定がないと直前の値が使われる
			Field5 = iota     // constの5行目なのでiotaは4になる
			Field6 = iota     // constの6行目なのでiotaは5になる
			Field7            // 値の指定がないと直前の値 (つまりiota) が使われる -> 6
		)
		fmt.Println(Field1, Field2, Field3, Field4, Field5, Field6, Field7) // 0 2 20 20 4 5 6

		// 「賢い」方法であるが，コメントを入れるべきである
		type BitField int
		const (
			BitField1 BitField = 1 << iota // 1になる
			BitField2                      // 2になる
			BitField3                      // 4になる
			BitField4                      // 8になる
		)
		fmt.Println(BitField1, BitField2, BitField3, BitField4) // 1 2 4 8
	}
	// 埋め込みによる合成
	//   - GoF以来，クラス継承よりオブジェクト合成のほうがよいとされている
	//   - Goには継承はなく，合成や昇格でコードを再利用する
	//   - 埋め込みフィールドを用いて，上位の構造体に昇格させる
	//   - 構造体に埋め込めるのは，構造体に限らずどんな型でも埋め込める
	//   - 埋め込まれた型のメソッドは，埋め込んでいる構造体のメソッドに昇格する
	example005()
	// 埋め込みで同名のフィールドやメソッドがある場合は，埋め込まれているほうが隠されるので，フィールドの型を明示してアクセスする必要がある
	{
		type Inner struct {
			X int
		}
		type Outer struct {
			Inner
			X int
		}

		o := Outer{
			Inner: Inner{
				X: 10,
			},
			X: 20,
		}
		fmt.Println(o.X)       // 20
		fmt.Println(o.Inner.X) // 10
	}
	// 埋め込みは継承ではない
	{
		type Employee struct {
			Name string
			ID   string
		}
		type Manager struct {
			Employee
			Reports []Employee
		}
		m := Manager{
			Employee: Employee{
				Name: "上杉謙信",
				ID:   "12345",
			},
			Reports: []Employee{},
		}
		// Employeeのフィールドに明示的にアクセスする
		var eOK Employee = m.Employee // OK!
		fmt.Println(eOK)              // {上杉謙信 12345}
		// Employee型の変数にManager型の値を代入することはできない
		// var eFail Employee = m        // コンパイル時のエラー！
		// cannot use m (variable of struct type Manager) as Employee value in variable declaration
	}
	// Goの具象型には動的ディスパッチがない
	//   - 動的ディスパッチ: 実行時の型によって呼び出し先 (メッセージのディスパッチ先) を決定すること
	example006()
	// インタフェース
	//   - インタフェースリテラルは interface { メソッドセットのシグネチャ } で定義する
	//   - メソッドセットとはインタフェースで定義されるメソッドの集まり
	//   - Goのインタフェースは暗黙的に実装される
	//     - すなわち，メソッドセットが具象型に完全に含まれていればインタフェースを実装したことに自動的になる
	// インタフェースは型安全なダックタイピング
	//   - Goはダックタイピングの柔軟さもインタフェースの明瞭さも両方必要だと考えた
	//   - Goのインタフェースは呼び出し側が必要なものを指定する
	//   - インタフェースIを満たす型に，Iにないメソッドを定義しても良い
	example007()
	// 標準のインタフェースを用いてデコレータパターンで実装する
	{
		process := func(r io.Reader) error { return nil }

		read := func(fileName string) error {
			r, err := os.Open(fileName)
			if err != nil {
				return err
			}
			defer r.Close()
			return process(r)
		}

		readCompressed := func(fileName string) error {
			r, err := os.Open(fileName)
			if err != nil {
				return err
			}
			defer r.Close()
			gz, err := gzip.NewReader(r)
			if err != nil {
				return err
			}
			defer gz.Close()
			return process(gz)
		}

		fmt.Println(read, readCompressed) // 0x1005d84e0 0x1005d82a0
	}
	// 構造体に型を埋め込めるのと同様に，インタフェースにインタフェースを埋め込むこともできる
	{
		type Reader interface {
			Read(p []byte) (n int, err error)
		}
		type Closer interface {
			Close() error
		}
		type ReadCloser interface {
			Reader // インタフェースReaderを埋め込む
			Closer // インタフェースCloserを埋め込む
		}
	}
	// 「インタフェースを受け取り，構造体を返す」
	//   - Goの関数やメソッドはインタフェースを受け取るべきである
	//     - コードが柔軟になり，どのような機能を使われているかが正確かつ明示的に宣言される
	//     - 例外的にGCによるパフォーマンス低下がある場合は構造体を受け取るよう変更することもある
	//   - Goの関数やメソッドは構造体を返すべきである
	//     - 新しいバージョンのコードで関数の戻り値を徐々に更新しやすい
	//     - インタフェースを返すと，そのインタフェースを実装する全ての既存実装を修正する必要がある
	//     - 具象型なら，既存コードを破壊することなくメソッドやフィールドを追加できる
	//   - 例外的にインタフェースを返す必要があることもある
	//   - 基本的に各具象型に対して各々のファクトリ関数を用意すべきである
	//     - ただし，エラーは error インタフェースを返さざるを得ない
	// インタフェースとnil
	//   - Goのインタフェースは値と値の型の2つのポインタフィールドをもつ構造体である
	//     - 型フィールドが非nilならば，インタフェースも非nilである
	//     - 値を指すポインタが非nilならば，インタフェースも非nilである (型のない変数はない)
	//     - インタフェースがnilであるとは，型フィールドも値を指すポインタも両方ともnilであることを意味する
	example008()
	// インタフェース型の変数iがnilかどうかは，メソッドを起動できるかどうかを示すことになる
	//   - インタフェース型の値のポインタがnilでも，ポインタレシーバでかつnilを処理できるメソッドならば問題ない
	//   - 値型レシーバやnilを処理できなければパニックになる
	//   - 非nil型のインタフェースに結びつけられた値がnilかどうかは，すぐには判断がつかずリフレクションを利用する必要がある
	// インタフェースは比較可能である
	//   - 前述の通り，インタフェースがnilと等しいとは，型と値のフィールドの両方がnilである場合
	//   - 2つのインタフェース型が等しいとは，型も値も等しい場合である
	//   - 型が比較可能でない場合は実行時パニックになる可能性がある
	//   - マップのキーは比較可能でなければならないから，インタフェースを指定することもできる
	//     - この際にキーが比較可能でなければ，やはりパニックになる
	//     - リフレクションを利用して，比較を行う前に reflect.Value のメソッド Comparable で確認できる
	example009()
	// 空インタフェース: 他言語のanyを表すために interface{} 「空(くう)インタフェース」を利用できる
	//   - Go 1.18以降はanyが使えるため，そちらを利用する
	//   - 空インタフェースは特殊な構文ではなく，満たすべきメソッドが0個のインタフェースである
	//   - Goは強い型付け言語なので，基本的にanyを使うことは避けるべきである
	//     - any を読み出すためにリフレクションや型アサーション，型switchを利用する
	{
		var i any // 「var i interface{}」も可
		i = 20
		fmt.Println(i) // 20
		i = "hello"
		fmt.Println(i) // hello
		i = struct {
			FirstName string
			LastName  string
		}{"信玄", "武田"}
		fmt.Println(i) // {信玄 武田}
	}
	// JSONファイルのような形式不明な外部ソースのプレースホルダーとしても使われる
	{
		readJson := func(fileName string) (map[string]any, error) {
			data := map[string]any{} // string->anyのマップで要素なし
			contents, err := os.ReadFile(fileName)
			if err != nil {
				return nil, err
			}
			json.Unmarshal(contents, &data)
			return data, nil
		}
		data, err := readJson("sample.json")
		if err == nil {
			fmt.Println(data)
		}
	}
	// 型アサーション: インタフェース型の値を，具象型だと仮定してアクセスする
	//   - 型アサーションと型変換は異なる (型アサーションは実行時，型変換はコンパイル時にチェックされる)
	//   - 型アサーションを行う際は必ず「カンマ ok イディオム」を利用すること
	{
		type MyInt int
		var i any // iには任意の型の値を記憶できる
		var mine MyInt = 20
		i = mine
		i2 := i.(MyInt)     // 型アサーション。iをMyInt型だと仮定 (assert) そて値 (20) をもらう
		fmt.Println(i2 + 1) // 21
		// 型アサーションが間違っているとパニックになる
		// i2 := i.(string)   // panic: interface conversion: interface {} is main.MyInt, not string
		// 厳密に値の型が一致していないとパニックになる (基底型が同じでもダメ)
		// i3 := i.(int)      // panic: interface conversion: interface {} is main.MyInt, not int
		// パニックを起こさないためには，「カンマ ok イディオム」を利用する
		i3, ok := i.(int) // 失敗する型アサーション
		if !ok {
			err := fmt.Errorf("iの型 (値:%v) が想定外です", i)
			fmt.Println(err.Error()) // エラーメッセージを表示
			// os.Exit(1)            // プログラムを終了
		} else {
			fmt.Println(i3)
		}
	}
	// 型switch: インタフェースが複数の型のいずれかである可能性がある場合に用いる
	//   - インタフェースiに対して switch i.(type) { case 型: ... } とする
	//   - イディオム的には型を特定した変数を同名でシャドーイングして i := i.(type) とする
	{
		type MyInt int
		doThings := func(i any) {
			switch j := i.(type) {
			case nil: // iはnil。jの型はany
				fmt.Printf("  case nil; i:%v (型:%T), j:%v (型:%T)\n", i, i, j, j)
			case int: // jの型はint
				fmt.Printf("  case int; i:%d (型:%T), j:%v (型:%T)\n", i, i, j, j)
			case MyInt: // jの型はMyInt
				fmt.Printf("  case MyInt; i:%d (型:%T), j:%v (型:%T)\n", i, i, j, j)
			case io.Reader: // jの型はio.Reader
				fmt.Printf("  case io.Reader; i:%v (型:%T), j:%v (型:%T)\n", i, i, j, j)
			case string: // jは文字列
				fmt.Printf("  case string; i:%s (型:%T), j:%v (型:%T)\n", i, i, j, j)
			case bool, rune: // iはboolかruneなので，jの型はany
				fmt.Printf("  case string; i:%v (型:%T), j:%v (型:%T)\n", i, i, j, j)
			default: // iの型は不明。jの型はany
				fmt.Printf("  default; i:%v (型:%T), j:%v (型:%T)\n", i, i, j, j)
			}
		}
		doThings(nil)
		doThings(3)
		doThings("abc")
		doThings(struct{}{})
	}
	// 型アサーションや型switchの利用は頻繁に使うべきではない
	//   - 関数が処理をするために必要な型を正確に宣言していないことになる
	//   - 型アサーションが利用される場合の例
	//     - あるインタフェースとして受け取った具象型が別のインタフェースを実装しているかを確認する
	//     - コンテキストの追加などでAPIを破壊的ではない方法で更新する
	// 型アサーションや型switchではデコレータパターンでラップされたオプションのインタフェースを検知できない
	example010()
	// 関数型とインタフェース
	//   - Goは任意のユーザー定義型にメソッドを追加できるため，関数型にもメソッドを追加できる
	//   - よって関数がインタフェースを実装できることになる
	// 暗黙のインタフェースによる依存性注入
	//   - 暗黙のインタフェースを利用すると依存性注入を用意に実装できる (追加ライブラリも不要)
	//   - インタフェースを満たしていれば，具象型を入れ替えられる
	example011()
	// Goは強くオブジェクト指向的ではないが，オブジェクト指向言語や関数型言語の要素を取り入れて，
	// シンプルで可読性が高く，大規模なプロジェクトで長期間使えるメンテナンスしやすい言語を目指している
}
