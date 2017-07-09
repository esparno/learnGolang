package main

import (
	"net/http"
	"io"
	"os"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/alister.jpg", dogPic)
	http.ListenAndServe(":8080", nil)
}
func index(wr http.ResponseWriter, req *http.Request){
io.WriteString(wr, "Index")
}
func dog(wr http.ResponseWriter, req *http.Request){
	wr.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(wr, `<img src="/alister.jpg" height="25%" width="25%">`)
}
func dogPic(wr http.ResponseWriter, req *http.Request){
	f, err := os.Open("alister.jpg")
	if err != nil {
		http.Error(wr, "File not found", 404)
		return
	}
	defer f.Close()
	io.Copy(wr, f)
}