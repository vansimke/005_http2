package main

import (
	"net/http"

	"golang.org/x/net/http2"
)

func main() {
	server := new(http.Server)
	h2Conf := new(http2.Server)

	http2.ConfigureServer(server, h2Conf)

	http.Handle("/", http.FileServer(http.Dir("public")))
	server.Addr = ":3001"
	server.ListenAndServeTLS("cert.pem", "key.pem")
}
