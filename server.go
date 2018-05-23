package main

import (
	"go-webserver/logger"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	logger.Info("hello world request")
	w.Write([]byte("Hello World!!"))
}

func main() {
	http.HandleFunc("/", hello)

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		logger.Panic("ListenAndServe Panic! ", err)
	}
}
