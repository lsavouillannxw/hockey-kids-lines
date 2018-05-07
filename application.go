package main

import (
	"net/http"
	"HockeyLines/rest"
)

func init() {
	http.HandleFunc("/", rest.Handler)
}
