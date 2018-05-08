package main

import (
	"github.com/lsavouillannxw/hockey-kids-lines/rest"
	"net/http"
)

func init() {
	http.HandleFunc("/", rest.Handler)
}
