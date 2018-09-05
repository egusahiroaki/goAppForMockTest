package main

import (
        "fmt"
        "time"
)

var timeNowFunc = time.Now // 関数をグローバル変数に代入

func main() {
        fmt.Println(Greet("マトペ"))
}

func Greet(n string) string {
        t := timeNowFunc() // グローバル変数でtime.Now関数を呼び出し
        if 6 <= t.Hour() && t.Hour() <= 18 {
                return fmt.Sprintf("こんにちは%sさん。今は%d時ですよ！", n, t.Hour())
        } else {
                return fmt.Sprintf("こんばんは%sさん。今は%d時ですよ！", n, t.Hour())
        }
}