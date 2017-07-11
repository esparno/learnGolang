package main

import (
	"html/template"
	"net/http"
	"github.com/satori/go.uuid"
	"fmt"
)

var tpl *template.Template
var Users = map[string]user{}
var Sessions = map[string] string{}
func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
func index(wr http.ResponseWriter, req *http.Request){
	cookie, err := req.Cookie("session")
	if err != nil {
		cookie = &http.Cookie {
			Name: "session",
			Value: uuid.NewV4().String(),
		}
		http.SetCookie(wr, cookie)
	}

	var u user
	if un, ok := Sessions[cookie.Value]; ok {
		u = Users[un]
	}

	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		fn := req.FormValue("firstname")
		ln := req.FormValue("lastname")
		u = user{ UserName:un, FirstName: fn, LastName: ln}
		Sessions[cookie.Value] = u.UserName
		Users[u.UserName] = u
	}
	tpl.ExecuteTemplate(wr, "index.gohtml", u)
	fmt.Println(u)
}
func bar (wr http.ResponseWriter, req *http.Request){
	cookie, err := req.Cookie("session")
	if err != nil {
		fmt.Println("here")
		http.Redirect(wr, req, "/", http.StatusSeeOther)
		return
	}
	un, ok := Sessions[cookie.Value]
	if !ok {
		http.Redirect(wr, req, "/", http.StatusSeeOther)
		return
	}
	u := Users[un]
	tpl.ExecuteTemplate(wr, "bar.gohtml", u)
}

type user struct {
	UserName string
	FirstName string
	LastName string
}