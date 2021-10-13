package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/healthz", healthz)
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/header", modifyResHeader)

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// 1.接收客户端 request，并将 request 中带的 header 写入 response header
func modifyResHeader(w http.ResponseWriter, r *http.Request) {
	for i, v := range r.Header {
		for _, v2 := range v {
			w.Header().Add(i, v2)
		}
	}

	// 2.读取当前系统的环境变量中的 VERSION 配置，并写入 response header
	version := os.Getenv("VERSION")
	w.Header().Add("Version", version)

	resHeader, err := json.Marshal(w.Header())
	if err != nil {
		log.Fatal(err)
	}

	_, err = w.Write(resHeader)
	if err != nil {
		log.Fatal(err)
	}
}

// 4.当访问 localhost/healthz 时，应返回 200
func healthz(w http.ResponseWriter, r *http.Request) {
	if r.Host == "localhost" || r.Host == "127.0.0.1" {
		_, err := w.Write([]byte("200"))
		if err != nil {
			log.Fatal(err)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("hello golang"))
	if err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(http.StatusUnauthorized)
}
