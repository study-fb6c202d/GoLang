package main

import (
	"fmt"
	"log"
	"net/http"
)

import (
	"WebServer/middleware"
	"WebServer/router"
)

const PORT int = 9090

func main() {
	http.HandleFunc("/", middleware.Logger(router.Index))
	err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil)
	if err != nil {
		log.Fatal("ListenAndServe Error", err)
	} else {
		fmt.Printf("ListenAndServe Started! -> Port(%d)", PORT)
	}
}
