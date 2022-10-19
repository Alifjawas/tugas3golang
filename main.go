package main

import (
	"net/http"
	"time"
	"tugas-3-gin/handler"
	"tugas-3-gin/helper"
	"tugas-3-gin/service"
)

func main() {
	go autoUpdate()
	http.HandleFunc("/", handler.Handler)

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}
}

func autoUpdate() {
	disaster := helper.GetDisaster()
	for range time.Tick(15 * time.Second) {
		service.UpdateJSON(disaster)
	}
}
