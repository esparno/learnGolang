package main

import (
	"net/http"
	"io"
)

func (d Dog)ServeHTTP(wr http.ResponseWriter, req *http.Request){
	io.WriteString(wr, "Woof!")
}
func main() {
	var d Dog

	http.Handle("/dog/", d)
	http.ListenAndServe(":8080", nil)
}
// because of the / after dog, all routes beginning with /dog/ will lead to handled function
type Dog struct {
	Name string
	Breed string
	Age int
}