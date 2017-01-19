package main

import (
	"compufour/pkg/service"
	"net/http"
)

func main() {
	http.HandleFunc("/", service.HandleIndex)
	http.ListenAndServe(":8082", nil)
}
