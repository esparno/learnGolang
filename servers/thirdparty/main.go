package main

import (
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
	"log"
)

var tpl *template.Template
func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}
func main() {
	mux := httprouter.New()
	mux.GET("/", index)
	mux.GET("/about", about)
	mux.GET("/contact", contact)
	mux.GET("/apply", apply)
	mux.POST("/apply", applyPost)
	http.ListenAndServe(":8080", mux)
}
func HandleError(wr http.ResponseWriter, err error){
	if err != nil {
		http.Error(wr, err.Error(), http.StatusInternalServerError )
		log.Fatalln(err)
	}
}
func index(wr http.ResponseWriter, req *http.Request, _ httprouter.Params){
	err := tpl.ExecuteTemplate(wr, "index.gohtml", nil)
	HandleError(wr, err)
}
func about(wr http.ResponseWriter, req *http.Request, _ httprouter.Params){
	err := tpl.ExecuteTemplate(wr, "about.gohtml", nil)
	HandleError(wr, err)
}
func contact(wr http.ResponseWriter, req *http.Request, _ httprouter.Params){
	err := tpl.ExecuteTemplate(wr, "contact.gohtml", nil)
	HandleError(wr, err)
}
func apply(wr http.ResponseWriter, req *http.Request, _ httprouter.Params){
	err := tpl.ExecuteTemplate(wr, "apply.gohtml", nil)
	HandleError(wr, err)
}
func applyPost(wr http.ResponseWriter, req *http.Request, _ httprouter.Params){
	err := tpl.ExecuteTemplate(wr, "applypost.gohtml", nil)
	HandleError(wr, err)
}