package main

import (
	"athgene/twentythreeandme"
	"net/http"
)

func main() {
	http.HandleFunc("/", twentythreeandme.HandleDownload)
	http.ListenAndServe(":8080", nil)
}
