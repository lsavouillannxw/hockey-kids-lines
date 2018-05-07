package main

import (
	"net/http"
	"github.com/lsavouillannxw/hockey-kids-lines/rest"
)

func init() {
	http.HandleFunc("/", rest.Handler)
}
