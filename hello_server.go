package main

import (
	"fmt"
	"net/http"
)

func main() {

	// 「/a」に対して処理を追加
	http.HandleFunc("/a", handler)

	// 8080ポートで起動
	// 第一引数が""の場合はデフォルトの80番が利用される
	// 第二引数がnilの場合は、デフォルトのマルチプレクサであるDefaultServeMuxが使われる
	http.ListenAndServe(":8080", nil)
}

// リクエストを処理する関数
func handler(writer http.ResponseWriter, Request *http.Request) {
	// FPrintはPrint（標準出力）接頭辞としてFがついている
	// 書き込み先を明示的に指定できる
	fmt.Fprint(writer, "Hello World from Go.")
}
