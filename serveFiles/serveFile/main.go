package main

import (
	"net/http"
	"io"
	"os"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/dog/", dog)
	http.ListenAndServe(":8080", nil)
}

func dog(wr http.ResponseWriter, req *http.Request){
	wr.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(wr, `<img src="/alister.jpg" height="25%" width="25%">`)
}
