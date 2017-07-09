package main

import (
	"net/http"
	"io"
)

func main() {
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./assets"))))
	http.HandleFunc("/dog/", dog)
	http.ListenAndServe(":8080", nil)
}
func dog (wr http.ResponseWriter, req *http.Request) {
	wr.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(wr, `<img src="/resources/alister.jpg" height="25%" width="25%">`)
}
