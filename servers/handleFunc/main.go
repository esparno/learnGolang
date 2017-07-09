package main

import (
	"net/http"
	"io"
)

func Woof(wr http.ResponseWriter, req *http.Request){
	io.WriteString(wr, "Woof!")
}
func main() {
	http.HandleFunc("/dog/", Woof)
	http.ListenAndServe(":8080", nil)
}