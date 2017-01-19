package main

import (
	"net/http"
	"compufour/pkg/service"
)

func main() {
	http.HandleFunc("/", service.HandleIndex)
	http.ListenAndServe(":8082", nil)
}
