/**
 * 2章 事前宣言された型
 */

package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	// 事前宣言された型: 整数・浮動小数点数・文字列・論理知などのいわゆる「組み込み型」
	// ゼロ値: 変数が宣言されたが値が割り当てられていない場合に自動的に割り当てられる値

	// リテラル: 数値や文字，文字列など特定のデータ型の値をソースコード内に直接書いたもの
	// 整数リテラル: 0b1010 (2進数), 0o644 (8進数), 42 (10進数), 0x2A (16進数),
	//    桁区切りのアンダースコア: 1_000_000
	fmt.Println(0b1010, 0o644, 42, 0x2A, 1_000_000)
	// 浮動小数点リテラル: 3.14, 2.5e10, 1.0e-5, 桁区切りのアンダースコア: 1_000.000_1
	fmt.Println(3.14, 2.5e10, 1.0e-5, 1_000.000_1)
	// rune (ルーン) リテラル: (文字を表す) 'a' == '\141' == '\x61' == '\u0061' == '\U00000061',
	//    'あ’ (Unicode 1文字), '\n' (改行などバックスラッシュでエスケープされた文字)
	fmt.Println(string('a'), string('\141'), string('\x61'), string('\u0061'), string('\U00000061'))
	// 文字列リテラル:
	//   - 解釈対象の文字列リテラル: "Hello, World!",
	//   - ロー文字列リテラル: `Raw \string literal<<改行>>Next line` (<<改行>>は '\n' ではなく改行そのもの)
	// リテラルは型付けされない，決定的でないときはデフォルトの型が利用される
	fmt.Println("Hello, World!")
	fmt.Println(`Raw string literal
Next line`)

	// 具体的な型
	// bool (論理型): true, false のいずれか。ゼロ値は false
	{
		var flag bool // false がゼロ値として割り当てられる
		isAwesome := true
		fmt.Println(flag, isAwesome)
	}
	// 整数型: int8, int16, int32, int64, uint8, uint16, uint32, uint64,
	//   int (CPU依存 == int32 or int64), uint (符号無し，CPU依存 == uint32 or uint64),
	//   rune (文字型, == int32), uintptr (ポインタ型), byte (== uint8)
	//   * 原則はint
	//   * 整数を対象とする場合はジェネリクスを利用する
	//   * バイナリファイルやネットワークストリームなどはそれぞれ具体的な型を利用する
	{
		var x int = 10
		x *= 2
		fmt.Println(x)
	}
	// 浮動小数点型: float32, float64
	//   * 非0の浮動小数点数を0で悪と +Inf / -Inf になる
	//   * 浮動小数点数を比較する際に == / != を使うのは避ける (イプシロン範囲内にあるかを判定)
	// 複素数型: complex64, complex128
	//   * デフォルトの型はcomplex128
	//   * complex(real, imag) の引数のうちいずれもfloat32型ならcomplex64になる
	{
		var x, y = complex(2.5, 3.1), complex(10.2, 2)
		fmt.Println(x+y, x-y, x*y, x/y, real(x), imag(x), cmplx.Abs(x))
	}
	// 文字型, 文字列型: rune, string
	//  * 文字列の結合には + 演算子を利用する
	//  * 文字であることを明示的に表すために極力 rune 型を利用すること
	{
		var myFirstInitial rune = 'J'
		var myLastInitial int32 = 'B'
		fmt.Println(string(myFirstInitial), string(myLastInitial), string(myFirstInitial)+"."+string(myLastInitial))
	}

	// 明示的型変換: Goでは暗黙の型変換を行わない，必ず明示的に型変換を行う
	// 特に bool 型の暗黙の型変換が行われないことに注意する
	{
		var x int = 10
		var y float64 = 30.2
		var sum1 float64 = float64(x) + y
		var sum2 int = x + int(y)
		fmt.Println(sum1, sum2)

		var b byte = 100
		var sum3 int = x + int(b)
		var sum4 byte = byte(x) + b
		fmt.Println(sum3, sum4)
	}

	// リテラルと型: リテラルは型付けされない
	// リテラル同士の演算ならば互換性があれば，許容される
	{
		var x float64 = 200.3 * 5
		var y float64 = 10
		var z = x * 3
		// ↓はコンパイルエラーになる (リテラル1000はbyte型に収まらない)
		// var b: byte = 1000
		fmt.Println(x, y, z)
	}

	// 変数宣言
	// varによる宣言: var 変数名 型 = 初期値
	//   - 省略形1: var 変数名 = 初期値 (型は初期値から推論される)
	//   - 省略形2: var 変数名 型 (ゼロ値が割り当てられる)
	{
		var x int = 10
		var y int
		fmt.Println(x, y)
	}
	// 多重代入: var 変数名1, 変数名2 型 = 初期値1, 初期値2
	//    - 省略形1: var 変数名1, 変数名2 型 (ゼロ値が割り当てられる)
	//    - 省略形2: var 変数名1, 変数名2 = 初期値1, 初期値2 (型は初期値から推論される, この場合は各変数が異なる型でも許容される)
	{
		var x1, y1 int = 10, 20
		fmt.Println(x1, y1)
		var x2, y2 int
		fmt.Println(x2, y2)
		var x3, y3 = 10, "hello"
		fmt.Println(x3, y3)
	}
	// 宣言リスト: 複数の変数の宣言をまとめることができる
	{
		var (
			x    int
			y        = 20
			z    int = 30
			d, e     = 40, "hello"
			f, g string
		)
		fmt.Println(x, y, z, d, e, f, g)
	}
	// :=による宣言: 変数名 := 初期値 (関数内でのみ利用可能，型は初期値から推論される)
	//   * := は既存の変数に値を代入することができる (一番内側のブロックで宣言されたものに限る)
	//   * := は関数の外では使えない
	{
		x := 10
		x, y := 30, "hello"
		fmt.Println(x, y)
	}
	// 変数宣言の選び方
	//   1. 関数内では := を使うのが基本
	//   2. 関数の外側でパッケージレベルの変数を宣言する際は宣言リストを利用する (ただ，これ自体は非推奨)
	//      * パッケージブロックではイミュータブルな変数のみを定義するべきである
	//   3. 関数内でも := を避けるべき場合
	//      3-1. ゼロ値を明示的に割り当てたい場合は var x int で定義する
	//      3-2. 型が指定されていないリテラル・定数をデフォルトでない型として利用する場合は var pi float32 = 3.14 で定義する
	//   4. var や := を利用して複数の変数宣言を1行で行うのは，複数の値を返す関数の戻り値を代入するときにするべきである

	// 定数: 値がイミュータブルであることを保証できる
	// Goの定数はリテラルに名前をつける機能しかない (コンパイル時に決定できる値でないといけない)
	//   - 数値リテラル
	//   - 文字列
	//   - rune
	//   - true/false.
	//   - 組み込み関数 complex, real, imag, len, cap の結果
	//   - 以上と演算子から構成される式
	const_function()

	// Goでは「変数」がイミュータブルであることを指定する方法は存在しない
	{
		x := 5
		y := 10
		// ↓ はコンパイルできない
		// const z = x + y // x + y (value of type int) is not constant
		var z = x + y
		fmt.Println(z)
	}

	// 型付きの定数・型のない定数
	{
		// 型のない定数
		const x = 10

		// 型付きの変数への代入は許容される
		var y int = x
		var z float64 = x
		var d byte = x
		fmt.Println(y, z, d)

		// 型付きの定数
		const typedX int = 10

		// 異なる型の型付きの変数への代入は許容されない
		// ↓ はコンパイルできない
		// var w float64 = typedX // cannot use typedX (constant 10 of type int) as float64 value in variable declaration
	}

	// 未使用変数
	// Goでは宣言された局所変数 (ローカル変数) はすべて使われなければならない
	// 宣言はされているが，使われていない変数があるとコンパイルエラーになる
	// ただし，厳密さはさほどなく，その変数が一度でも使われていればよい
	// (定数は未使用でもコンパイルエラーにならない)
	{
		x := 10 // このxの値は読まれることはない
		x = 20
		fmt.Println(x)
		x = 30 // このxの値も読まれることはない
	}

	// 変数と定数の命名
	// 変数および定数の名前は文字またはアンダースコアで始まり，続く文字は文字，数字，アンダースコアのいずれかである
	// ただし，文字や数字はUnicodeで文字や数字としてみなされているかである
	{
		_0 := 0_0
		_𝟙 := 20
		π := 3
		ａ := "hello"
		数値１ := 33
		fmt.Println(_0, _𝟙, π, ａ, 数値１)
	}
	{
		ａ := "こんにちは" // 全角の「ａ」
		a := "サヨウナラ" // 半角の「a」
		fmt.Println(ａ, a)
	}
	// Goではイディオム的にキャメルケース (camelCase) が利用される
	//   * スネークケース (SNAKE_CASE) は利用されない
	//   * パッケージレベルで外部公開されるか否かは，名前の先頭が大文字か否かで決まる
	// 関数内では基本的に短い名前を使う
	//   * for-rangeループではkeyにkを，valueにvを使うのが慣例
	//   * forループならばi, jなどがよく使われる
	//   * 慣用的に整数はi, 浮動小数点数はf, 論理値はbなどが使われる
}

// 定数 -------------------------------------------------------------------------

const x int64 = 10
const (
	idKey   = "id"
	nameKey = "name"
)
const z = 20 * 10

func const_function() {
	const y = "hello"

	fmt.Println(x)
	fmt.Println(y)

	// ↓ はコンパイルできない
	// y = x + 1 // cannot assign to y (neither addressable nor a map index expression)
	// ↓ はコンパイルできない
	// y = "bye" // cannot assign to y (neither addressable nor a map index expression)

	fmt.Println(x)
	fmt.Println(y)
}

// -----------------------------------------------------------------------------
