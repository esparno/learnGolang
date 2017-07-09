package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template
func init () {
	tpl, _ = template.ParseFiles("index.gohtml")
}
func main() {
	var Alister Dog
	http.ListenAndServe(":8080", Alister)
}
type Dog struct {
	Name string
	Breed string
	Age int
}
func (d Dog) ServeHTTP(wr http.ResponseWriter, req *http.Request) {
	wr.Header().Set("ERIN-KEY", "ALISTER")
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	tpl.ExecuteTemplate(wr, "index.gohtml", req.Form)
}
