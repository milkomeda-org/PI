// Copyright The ZHIYUN Co. All rights reserved.
// Created by vinson on 2020/10/20.

package main

import (
	"bytes"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// 主动关闭服务器
var server *http.Server

func main() {
	// 一个通知退出的chan
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	mux := http.NewServeMux()
	mux.HandleFunc("/", hand)

	server = &http.Server{
		Addr:         ":" + os.Getenv("PORT"),
		WriteTimeout: time.Second * 4,
		Handler:      mux,
	}

	go func() {
		// 接收退出信号
		<-quit
		if err := server.Close(); err != nil {
			log.Fatal("Close server:", err)
		}
	}()

	log.Println("Starting v3 httpserver")
	err := server.ListenAndServe()
	if err != nil {
		// 正常退出
		if err == http.ErrServerClosed {
			log.Fatal("Server closed under request")
		} else {
			log.Fatal("Server closed unexpected", err)
		}
	}
	log.Fatal("Server exited")

}

func hand(w http.ResponseWriter, r *http.Request) {
	w.Write(bytes.NewBufferString("hello pi").Bytes())
}
