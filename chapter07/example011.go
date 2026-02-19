package main

import (
	"errors"
	"fmt"
	"net/http"
)

/** ログを出力するユーティリティ関数 */

// ログを出力
func LogOutput(message string) {
	fmt.Println(message)
}

/** データを保存しておくための簡単な仕組み */

// 文字列に対して文字列を記憶するマップ
type SimpleDataStore struct {
	userData map[string]string
}

// SimpleDataStoreに付随するメソッド。userIDを受け取り，名前を返す
func (sds SimpleDataStore) UserNameForID(userID string) (string, bool) {
	name, ok := sds.userData[userID]
	return name, ok
}

// 新しいSimpleDataStoreを返すファクトリ関数
func NewSimpleDataStore() SimpleDataStore {
	return SimpleDataStore{
		userData: map[string]string{
			"1": "花子",
			"2": "太郎",
			"3": "パット",
		},
	}
}

/**
 * ビジネスロジックの作成: ユーザーを探して「こんにちは」を言う
 *
 * ビジネスロジックはログ記録やデータストアへの依存がある
 * 具体的な型LogOutputやSimpleDataStoreに依存させたくないので，インタフェースを記述する
 */

// インタフェースDataStoreの定義
// SimpleDataStoreを抽象化したもの
type DataStore interface {
	UserNameForID(userID string) (string, bool)
}

// インタフェースLoggerの定義
// LogOutputを抽象化したもの
type Logger interface {
	Log(message string)
}

// 型LoggerAdapterを定義
type LoggerAdapter func(message string)

// LoggerAdapterにL付随するメソッド
func (lg LoggerAdapter) Log(message string) {
	lg(message) // 受け取ったメッセージmessageを引数として自分を呼び出す
}

type SimpleLogic struct {
	l  Logger
	ds DataStore
}

func (sl SimpleLogic) SayHello(userID string) (string, error) {
	sl.l.Log("in SayHello for " + userID)
	name, ok := sl.ds.UserNameForID(userID)
	if !ok {
		return "", errors.New("未登録のユーザーです")
	}
	return name + "さん、こんにちは", nil
}

func NewSimpleLogic(l Logger, ds DataStore) SimpleLogic {
	return SimpleLogic{
		l:  l,
		ds: ds,
	}
}

type BusinessLogic interface {
	SayHello(userID string) (string, error)
}

/** MVCのコントローラとアクションメソッドの定義 */

type Controller struct {
	l     Logger
	Logic BusinessLogic
}

func (c Controller) SayHello(w http.ResponseWriter, r *http.Request) {
	c.l.Log("In SayHello")
	userID := r.URL.Query().Get("user_id")
	message, err := c.Logic.SayHello(userID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(message))
}

func NewController(l Logger, logic BusinessLogic) Controller {
	return Controller{
		l:     l,
		Logic: logic,
	}
}

func example011() {
	l := LoggerAdapter(LogOutput)
	ds := NewSimpleDataStore()
	logic := NewSimpleLogic(l, ds)
	c := NewController(l, logic)
	http.HandleFunc("/hello", c.SayHello)
	//http.ListenAndServe(":8080", nil)
}
