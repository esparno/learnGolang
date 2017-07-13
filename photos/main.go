package main

import (
	"html/template"
	"net/http"
	"fmt"
	"strings"
	"crypto/sha1"
	"io"
	"os"
	"path/filepath"
	"github.com/satori/go.uuid"
)

var tpl *template.Template
func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}
func main() {
	http.HandleFunc("/", index)
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("./public"))))
	http.ListenAndServe(":8080", nil)
}
func index (wr http.ResponseWriter, req *http.Request) {
	cookie := getSessionCookie(wr, req)
	if req.Method == http.MethodPost {
		multipartfile, fileheader, err:= req.FormFile("nf")
		if err != nil {
			fmt.Println(err)
		}
		defer multipartfile.Close()

		// grab file extension
		ext := strings.Split(fileheader.Filename, ".")[1]

		h := sha1.New()
		io.Copy(h, multipartfile)
		filename := fmt.Sprintf("%x.%s",h.Sum(nil),ext)

		wd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
		}
		path := filepath.Join(wd, "public", "pics", filename)
		nf, err := os.Create(path)
		if err != nil {
			fmt.Println(err)
		}
		defer nf.Close()
		multipartfile.Seek(0,0)
		io.Copy(nf, multipartfile)
		cookie = appendCookie(wr, cookie, filename)
	}
	xs := strings.Split(cookie.Value, "|")
	tpl.ExecuteTemplate(wr, "index.gohtml", xs[1:])
}
func getSessionCookie(wr http.ResponseWriter, req *http.Request) *http.Cookie{
	cookie, err := req.Cookie("session")
	if err != nil {
		cookie = &http.Cookie{
			Name: "session",
			Value: uuid.NewV4().String(),
		}
		http.SetCookie(wr, cookie)
	}
	return cookie
}
func appendCookie (wr http.ResponseWriter, cookie *http.Cookie, filename string) *http.Cookie {
	s := cookie.Value
	if !strings.Contains(s, filename) {
		s += "|" + filename
	}
	cookie.Value = s
	http.SetCookie(wr, cookie)
	return cookie
}