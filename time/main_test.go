// main_test.go
package main

import (
	"fmt"
	"testing"
	"time"
)

const timeformat = "2006-01-02 15:04:06" // timeのフォーマット指定文字列

// timeNowFuncグローバル変数の関数を入れ替える関数
func setNow(t time.Time) {
	timeNowFunc = func() time.Time {
		fmt.Println("[SetNow]")

		return t
	}
}

func TestMain(t *testing.T) {
	fmt.Println("[Test Start] TestUserServiceGetTotalPointByID")
	evening, _ := time.Parse(timeformat, "2014-08-14 14:10:00")
	night, _ := time.Parse(timeformat, "2014-08-14 22:30:00")

	// 昼のテスト
	setNow(evening) // スタブ！timeNowFuncに昼時刻を返す関数をセット
	if ret := Greet("まとぺ"); ret != "こんにちはまとぺさん。今は14時です！" {
		t.Errorf("Greet Fails. Got [%s]\n", ret)
	}

	// 夜のテスト
	setNow(night) // スタブ！timeNowFuncに昼時刻を返す関数をセット
	if ret := Greet("まとぺ"); ret != "こんばんはまとぺさん。今は22時ですよ！" {
		t.Errorf("Greet Fails. Got [%s]\n", ret)
	}
}
