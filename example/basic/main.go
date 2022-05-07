package main

import (
	"fmt"
	"net/http"

	"github.com/go-zoox/captcha"
)

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/captcha", captchaHandler)

	fmt.Println("Server start at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func rootHandler(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("hello, world"))
}

func captchaHandler(w http.ResponseWriter, _ *http.Request) {
	cap := captcha.New()
	fmt.Println("text:", cap.Text())
	cap.Write(w)
}
