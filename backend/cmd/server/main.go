package main

import (
	"net/http"

	"github.com/keithrooney/publik/internal"
)

func main() {
	http.Handle("/v1/echo", internal.NewDispatcher())
	http.ListenAndServe(":8080", nil)
}
