# Captcha - Powerful and easy to use

[![PkgGoDev](https://pkg.go.dev/badge/github.com/go-zoox/captcha)](https://pkg.go.dev/github.com/go-zoox/captcha)
[![Build Status](https://github.com/go-zoox/captcha/actions/workflows/ci.yml/badge.svg?branch=master)](https://github.com/go-zoox/captcha/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-zoox/captcha)](https://goreportcard.com/report/github.com/go-zoox/captcha)
[![Coverage Status](https://coveralls.io/repos/github/go-zoox/captcha/badge.svg?branch=master)](https://coveralls.io/github/go-zoox/captcha?branch=master)
[![GitHub issues](https://img.shields.io/github/issues/go-zoox/captcha.svg)](https://github.com/go-zoox/captcha/issues)
[![Release](https://img.shields.io/github/tag/go-zoox/captcha.svg?label=Release)](https://github.com/go-zoox/captcha/tags)

## Installation
To install the package, run:
```bash
go get github.com/go-zoox/captcha
```

## Getting Started

```go
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
```

## License
GoZoox is released under the [MIT License](./LICENSE).
