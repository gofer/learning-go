/**
 * 6章 ポインタ
 */

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
)

// T = string ならば func makePointer(s string) *string { return &s } と同じ意味
func makePointer[T any](v T) *T { // ヘルパー関数
	return &v
}

func main() {
	// ポインタ: ある値が保存されているメモリのアドレスを表す変数
	//   - ポインタのゼロ値は nil
	//   - C言語とは異なり int との相互変換はできない
	//   - Go にはガベージコレクションがある
	{
		var x int32 = 10
		var y bool = true
		pointerX := &x
		pointerY := &y
		var pointerZ *string

		fmt.Println(pointerX, pointerY, pointerZ) // 0x140000100f0 0x140000100f4 <nil>
	}
	// & はアドレス演算子で，変数の前に付けるとその変数のアドレスを返す (ポインタ型になる)
	{
		x := "hello"
		pointerToX := &x
		fmt.Println(pointerToX) // 0x14000120020
	}
	// * は関節参照演算子で，ポインタの前に付けるとそのポインタが指す変数の値を返す (デリファレンスとよぶ)
	{
		x := 10
		pointerToX := &x
		fmt.Println(pointerToX)  // 0x14000092028
		fmt.Println(*pointerToX) // 10
		z := 5 + *pointerToX
		fmt.Println(z) // 15
	}
	// nil ポインタをデリファレンスしようとするとパニックになる
	{
		var y *int
		fmt.Println(y == nil) // true
		// fmt.Println(*y)       // panic: runtime error: invalid memory address or nil pointer dereference
	}
	// 型の前に * を付けるとその型のポインタ型を表す
	{
		x := 10
		var pointerToX *int
		pointerToX = &x
		fmt.Println(pointerToX) // 0x14000010110
	}
	// 組み込み関数 new は指定された型のポインタ型変数を生成する (値はゼロ値で初期化される)
	// (ただし，new はあまり使われない)
	{
		var x = new(int)      // xの参照先にはintのゼロ値(0)が記憶される
		fmt.Println(x == nil) // false
		fmt.Println(x)        // 0x140000a0030 (例: xのアドレスが表示される)
		fmt.Println(*x == 0)  // true
		fmt.Println(*x)       // 0
	}
	// 構造体のポインタ型変数は構造体リテラルの前に & を付けて生成できる
	// (事前宣言された型のリテラルや定数の前に & を付けることはできない)
	{
		type Foo struct{} // 「Foo{}」は構造体リテラル
		x := &Foo{}
		var y string
		z := &y               // string へのポインタ
		fmt.Println(x == nil) // false
		fmt.Println(z == nil) // false
	}
	// 構造体のフィールドに事前宣言された型へのポインタがあると，そのフィールドにはリテラルを直接代入できない
	{
		type person struct {
			FirstName  string
			MiddleName *string
			LastName   string
		}

		p := person{
			FirstName: "Pat",
			// MiddleName: "Perry", // ← コンパイル時のエラー (cannot use "Perry" (untyped string constant) as *string value in struct literal)
			LastName: "Peterson",
		}
		fmt.Println(p) // {Pat <nil> Peterson}
	}
	// ↑の問題を回避する方法は2つある
	// 1. 定数値を保持する変数を作る
	{
		type person struct {
			FirstName  string
			MiddleName *string
			LastName   string
		}

		s := "Perry"
		p := person{
			FirstName:  "Pat",
			MiddleName: &s,
			LastName:   "Peterson",
		}
		fmt.Println(p) // {Pat 0x140000a4030 Peterson}
	}
	// 2. 値を受け取り，ポインタを返すジェネリックなヘルパー関数を作成する
	{

		type person struct {
			FirstName  string
			MiddleName *string
			LastName   string
		}
		p := person{
			FirstName:  "Pat",
			MiddleName: makePointer("Perry"), // これならうまくいく
			LastName:   "Peterson",
		}
		fmt.Println(p) // {Pat 0x140000a4040 Peterson}
	}
	// Go でイミュータブルとミュータブルの使い分けをするために値渡しとポインタ渡しを使い分ける
	//   - ポインタはミュータブルであることを宣言する手段として用いる
	//   - 関数にポインタを渡すと，関数はポインタのコピーを受け取る
	// 関数にポインタ型引数としてnilを渡すと，その値をnil以外に変えることはできない
	{
		failedUpdate := func(g *int) {
			x := 10
			g = &x
		}
		var f *int // f は nil
		failedUpdate(f)
		fmt.Println(f) // <nil>

		// 1. var f *int : 関数mainの変数fを定義し，nilで初期化する
		//   main.f (0x00000010) = 0
		// 2. failedUpdate(f) : 関数failedUpdate呼び出しで，変数fの値(nil)が引数gにコピーされる
		//   main.f (0x00000010) = 0
		//   failedUpdate.g (0x00001000) = 0
		// 3. x := 10 : 関数failedUpdateで変数xを定義し，10で初期化する
		//   main.f (0x00000010) = 0
		//   failedUpdate.g (0x00001000) = 0
		//   failedUpdate.x (0x00001004) = 10
		// 4. g = &x : 関数failedUpdateで引数gを変数xへのポインタにする
		//   main.f (0x00000010) = 0
		//   failedUpdate.g (0x00001000) = 0x00001004
		//   failedUpdate.x (0x00001004) = 10
		// 5. fmt.Println(f) : 関数mainの変数fの値はnilのままである
		//   main.f (0x00000010) = 0
	}
	// ポインタ型引数に代入された値が関数を終了しても消えずに残っていて欲しい場合は，デリファレンスして値を設定する必要がある
	{
		failedUpdate := func(px *int) {
			x2 := 20
			px = &x2
		}
		update := func(px *int) {
			*px = 20
		}
		x := 10
		failedUpdate(&x)
		fmt.Println(x) // 10
		update(&x)
		fmt.Println(x) // 20

		// 1. x := 10 : 関数mainで変数xを定義し，10で初期化する
		//   main.x (0x00000010) = 10
		// 2. failedUpdate(&x) : 関数failedUpdate呼び出しで，変数xへのポインタが引数pxにコピーされる
		//   main.x (0x00000010) = 10
		//   failedUpdate.px (0x00001000) = 0x00000010
		// 3. x2 := 20 : 関数failedUpdateで変数x2を定義し，20で初期化する
		//   main.x (0x00000010) = 10
		//   failedUpdate.px (0x00001000) = 0x00000010
		//   failedUpdate.x2 (0x00001004) = 20
		// 4. px = &x2 : 関数failedUpdateで引数pxを変数x2へのポインタにする
		//   main.x (0x00000010) = 10
		//   failedUpdate.px (0x00001000) = 0x00001004
		//   failedUpdate.x2 (0x00001004) = 20
		// 5. fmt.Println(x) : 関数mainの変数xの値は10のままである
		//   main.x (0x00000010) = 10
		// 6. update(&x) : 関数update呼び出しで，変数xへのポインタが引数pxにコピーされる
		//   main.x (0x00000010) = 10
		//   update.px (0x00002000) = 0x00000010
		// 7. *px = 20 : 関数updateで引数pxが指す変数の値を20にする
		//   main.x (0x00000010) = 20
		//   update.px (0x00002000) = 0x00000010
		// 8. fmt.Println(x) : 関数mainの変数xの値は20に変わっている
		//   main.x (0x00000010) = 20
	}
	// ポインタは最後の手段であり，利用に際して慎重に選択する必要がある
	//   - データフローがわかりづらい
	//   - ガベージコレクションの仕事が増える
	{
		type Foo struct {
			Field1 string
			Field2 int
		}

		// 悪い例: 関数にポインタを渡して中身を初期化している
		MakeFooBad := func(f *Foo) error {
			f.Field1 = "val"
			f.Field2 = 20
			return nil
		}

		var foo1 Foo
		_ = MakeFooBad(&foo1)
		fmt.Println("foo1:", foo1) // foo1: {val 20}

		// 良い例: 関数が構造体のインスタンスを生成して返す
		MakeFooGood := func() (Foo, error) {
			f := Foo{
				Field1: "val",
				Field2: 20,
			}
			return f, nil
		}

		foo2, _ := MakeFooGood()
		fmt.Println("foo2:", foo2) // foo2: {val 20}
	}
	// 関数がインターフェイスを受け取る際はポインタ引数を利用しなくてはならない
	//   - 特にJSONを扱う場合にこのパターンは頻出である
	//   - ただし，json.UnmarshalのAPIは例外的なものだと理解すること
	{
		f := struct {
			Name string // NameのNは大文字！小文字だと他パッケージから見えない
			Age  int
		}{}

		err := json.Unmarshal([]byte(`{"name": "小野小町", "occupation": "歌人", "age": 20}`), &f)
		// 大文字小文字の違いを無視して，フィールドに対応付けてくれる
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%+v\n", f) // {Name:小野小町 Age:20}
		// %v+ でフィールド名付きで出力
	}
	// サイズの大きい構造体を関数でやりとりするなどの場合はポインタの利用を検討する
	//   - 「値のコピーのコストが高い場合」ということ
	// 変数や構造体のフィールドが「ゼロ値」なのか「値がない」のかを区別する際ににもポインタを利用することがある
	//   - 「値がない」 = nil として表す
	//     - たとえばJSONのフィールドで null が許容されているときなど
	//   - 「ポインタを利用する」 = 「ミュータブルである」であることに注意する
	//     - 戻り値として nil に設定したポインタを戻すのでなく，「カンマ ok イディオム」で値とブール値を戻す
	//   - 引数として nil を渡す，構造体のフィールドに nil がある場合はそこに値を設定できないことに注意する
	// マップを関数の引数として渡すと，ポインタを渡すことは同じである
	//   - マップはポインタとして実装されている
	//   - したがって，外部公開のAPIの引数としてマップを利用することは望ましくない (型付け・GC処理の増加)
	//   - 可能な限り構造体の利用を優先する
	// スライスはメモリブロックへのポインタ，長さとキャパシティのint型フィールドの構造体である
	//   - コピー先のスライスの長さの変更は，コピー元のスライスからは見えない
	//     1. 関数mainにスライスdata が (array, len, cap) = (0x00001234, 3, 6) で存在する
	//       - main.data [0x00001000] = (0x00001234, 3, 6)
	//     2. 関数mainから関数fの引数xsにdataを渡す
	//       - main.data [0x00001000] = (0x00001234, 3, 6)
	//       - f.xs [0x00002000] = (0x00001234, 3, 6)
	//     3. 関数fでxsの長さを変更する (appendなど)
	//       - main.data [0x00001000] = (0x00001234, 3, 6)
	//       - f.xs [0x00002000] = (0x00001234, 4, 6)
	//     4. 関数mainに戻っても，元の長さは変わっていない
	//       - main.data [0x00001000] = (0x00001234, 3, 6)
	//   - キャパシティについても同様のことが言える
	//     - キャパシティを増やすと，データのストレージのポインタも変わってしまう
	//     - たとえば，もとのスライスが[1, 2, 3]で，4を末尾に追加して，全てを10倍して[10, 20, 30, 40]にしたとき，もとのスライスは[1, 2, 3]のままである
	// 関数の引数としてスライスを与えたとき，キャパシティが変わらなければ内容を変更できる
	//   - 再利用可能なバッファとして使える
	//   - データを読み込む度にメモリ割り当てを行うことを避けられる
	{
		processFile := func(fileName string) error {
			file, err := os.Open(fileName)
			if err != nil {
				return err
			}
			defer file.Close()
			data := make([]byte, 100)
			for {
				_, err := file.Read(data)
				// process(data[:count]) // 読み込んだデータの処理
				if err != nil {
					if errors.Is(err, io.EOF) {
						return nil
					}
					return err
				}
			}
		}

		err := processFile("main.go")
		if err != nil {
			fmt.Println("Error processing file:", err)
		}
	}
	// Goではコンパイル時に大きさが正確にわかる値 (事前宣言された型・配列・構造体) をスタックに，それ以外の値をヒープに割り当てる
	//   - ポインタが参照する変数をスタックに割り当てるためには，大きさがコンパイル時にわかる以外にも条件がある
	//     - その関数の戻り値でないこと (戻り値として受け取った後に参照先がなくなってしまう)
	//     - 別の関数の引数としてポインタを渡す場合でも同様の条件が成り立つこと
	//   - コンパイラがスタックにデータを割り当てできないと判断することを「エスケープした」という (エスケープ解析)
	//   - ヒープのメモリ管理を行うのがガベージコレクタである
	//     - GCの処理は重い，RAMへランダムアクセスする必要があるなどのパフォーマンス上の問題があり，ヒープへの割り当ては避けるべきものである
	//     - Goではなるべくスタックを利用し，GCの負荷を減らそうとしている
	// GCのチューニング
	//   - 環境変数 GOGC
	//     - GCは作業の最後にヒープサイズを確認し「(現在のヒープサイズ) * (1 + GOGC/100)」の値から次回のGCのトリガーを決定する
	//     - GOGC のデフォルト値は100である (つまり，ヒープサイズが2倍になるとGCがトリガーされる)
	//     - ざっくり，GOGC の値を2倍にする意図，GCのCPU使用率を半分になる
	//     - GOGC を off にするとGCが無効化されるが，避けるべき (メモリ不足になる)
	//  - 環境変数 GOMEMLIMIT
	//     - GOMEMLIMIT はGoアプリの総メモリ使用量を指定する (Javaの -Xmx オプションのようなもの)
	//     - GOMEMLIMIT のデフォルト値はmath.MaxInt64である (つまり，制限なし)
	//     - GOMEMLIMIT はバイト単位で指定できるが，B, KiB, MiB, GiBなどの単位も利用できる
	//     - メモリの最大使用量を制限すると，スラッシングを抑制できる
	//       - スラッシングとはヒープサイズが最大メモリ容量を超過してしまい，ディスクスワップが多く発生したために処理速度が著しく低下している状態のこと
	//     - Goはスラッシングの発生を検知すると，GCを終了して GOMEMLIMIT 以上のメモリを確保しようとする
}
