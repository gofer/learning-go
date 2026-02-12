package main

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"log"
	"os"
)

// defer: リソースのクリーンアップ処理を行う
//   - 他言語で言う try - finally のようなもの
//   - defer で指定された処理は関数終了時まで実行が延期される
//   - defer に指定できるのは関数 (メソッド・クロージャ)
//   - 複数の defer 文がある場合、LIFO (後入れ先出し) で実行される
//   - defer は return の後に実行される
//   - defer の引数は即時評価 (defer 文の場所で値が評価) される
//   - defer の戻り値を得る方法はない
//   - 名前付き戻り値を利用すると defer 内で戻り値を検証したり変更したりできる

func deferExample() {
	a := 10
	defer func(val int) {
		fmt.Println("first:", val)
	}(a)
	a = 20
	defer func(val int) {
		fmt.Println("second:", val)
	}(a)
	a = 30
	fmt.Println("exiting:", a)

	// exiting: 30
	// second: 20
	// first: 10
}

func DoSomeInserts(ctx context.Context, db *sql.DB, value1, value2 string) (err error) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() { // defer される関数の定義
		if err != nil {
			err = tx.Commit() // エラーがなければコミット
		}
		if err != nil {
			tx.Rollback() // コミットした結果エラーがあればロールバック
		}
	}() // 無名関数を実行

	_, err = tx.ExecContext(ctx, "INSERT INTO FOO (val) VALUES $1", value1)
	if err != nil {
		return err
	}
	// tx を使ってさらにデータベースに書き込むコードをここに追加する
	return nil
}

func getFile(name string) (*os.File, func(), error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, nil, err
	}
	return file, func() {
		file.Close()
	}, nil
}

func example008() {
	if len(os.Args) < 2 { // ファイル名が指定されているか
		log.Fatal("ファイルが指定されていません")
	}
	// f, err := os.Open(os.Args[1]) // ファイルをオープン
	f, closer, err := getFile(os.Args[1])
	if err != nil {
		log.Fatal(err) // オープンに問題あり。エラーを出力して終了
	}
	// defer f.Close() // 後始末のコード
	defer closer()

	data := make([]byte, 2048) // バイトのスライスを生成
	for {
		count, err := f.Read(data)    // 読み込んだバイト数とエラーを返す
		os.Stdout.Write(data[:count]) // 「標準出力」に出力
		if err != nil {
			if err != io.EOF { // ファイルの終わりでないならば
				log.Fatal(err) // エラーを出力して終了
			}
			break // forループを抜ける (ファイルの終わり)
		}
	}

	deferExample()
}
