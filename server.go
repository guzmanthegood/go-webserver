package main

import (
	"net/http"
)

import logger "github.com/guzmanweb/go-logger"

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func main() {
	http.HandleFunc("/", hello)

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		logger.Panic("ListenAndServe Panic! ", err)
	}
}
