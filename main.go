package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	counter "github.com/k-p5w/go-pixela-counter/api"
)

func main() {

	fmt.Println("start!")
	// // これで静的ファイルにアクセスできるとおもったのになあ
	// fs := http.FileServer(http.Dir("public"))
	// // http.Handle("/tool/", http.StripPrefix("/tool/", fs))

	// // `/tool/`以下のパスにアクセスされた場合、`fs`ハンドラーで処理
	// http.Handle("/", http.StripPrefix("/public/", fs))

	router := mux.NewRouter()
	router.HandleFunc("/api/ppppCounter", counter.HandlePPPCounter)

	err := http.ListenAndServe(":3000", router)
	if err != nil {
		fmt.Println("Failed to start server:", err)
		return
	}
	/*
		http.HandleFunc("/ppppCounter", counter.Handler)
		// http.HandleFunc("/view", widget.Handler)
		http.HandleFunc("/", counter.Handler)

		port := os.Getenv("PORT")
		if port == "" {
			port = "9999"
		}

		// 起動する
		http.ListenAndServe(":"+port, nil)
		// http.ListenAndServe("localhost:"+port, nil)
	*/
}
