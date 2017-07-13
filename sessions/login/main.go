package main

import (
	"html/template"
	"net/http"
	"fmt"
	"time"
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)


var tpl *template.Template
var users = make(map[string] *user)
var sessions = make (map[string] *session)
func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}
func main() {
	http.HandleFunc("/account", account)
	http.HandleFunc("/login", login)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}
func login(wr http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		cookie := getOrSetSession(wr, req)
		session := sessions[cookie.Value]
		un := req.FormValue("username")
		pw := req.FormValue("password")
		u, ok := users[un]
		if !ok {
			http.Error(wr,"Incorrect username/password", http.StatusForbidden)
			return
		}
		err := bcrypt.CompareHashAndPassword(u.Password, []byte(pw))
		if err != nil {
			http.Error(wr, "Incorrect username/password", http.StatusForbidden)
			return
		}
		session.username = un
		fmt.Println("login session", sessions[cookie.Value])
	}

	http.Redirect(wr, req, "/", http.StatusSeeOther)
}
func logout (wr http.ResponseWriter, req *http.Request){
	if req.Method == http.MethodPost {
		cookie := getOrSetSession(wr, req)
		cookie.MaxAge = -1
		http.SetCookie(wr, cookie)
	}
	http.Redirect(wr, req, "/", http.StatusSeeOther)
}
func signup(wr http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		fn := req.FormValue("firstname")
		ln := req.FormValue("lastname")
		pw := req.FormValue("password")
		if _, ok := users[un]; ok {
			http.Error(wr, "Username already taken", http.StatusForbidden)
			return
		}
		bs, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.MinCost)
		if err != nil {
			http.Error(wr, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		cookie := getOrSetSession(wr, req)
		session := sessions[cookie.Value]
		session.username = un
		users[un] = &user{
			Username: un,
			FirstName: fn,
			LastName: ln,
			Password: bs,
		}
	}
	http.Redirect(wr, req, "/", http.StatusSeeOther)
}
func index (wr http.ResponseWriter, req *http.Request){

	cookie := getOrSetSession(wr, req)
	cleanSessions()
	session := sessions[cookie.Value]
	u := users[session.username]
	tpl.ExecuteTemplate(wr, "index.gohtml", u)
}
func account(wr http.ResponseWriter, req *http.Request) {
	cookie := getOrSetSession(wr, req)
	session := sessions[cookie.Value]
	if session.username != ""{
		http.Redirect(wr, req, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(wr, "account.gohtml", nil)
}
func getOrSetSession(wr http.ResponseWriter, req *http.Request) *http.Cookie{
	cookie, err := req.Cookie("session")
	if err != nil {
		cookie = &http.Cookie{
			Name: "session",
			Value: uuid.NewV4().String(),
			MaxAge: 60,
		}
		http.SetCookie(wr, cookie)
		session := &session{ lastActivity: time.Now()}
		sessions[cookie.Value] = session
	}
	session := sessions[cookie.Value]
	session.lastActivity = time.Now()
	return cookie
}
func cleanSessions(){
	fmt.Println("BEFORE CLEAN")
	showSessions()
	for k, v := range sessions {
		if time.Now().Sub(v.lastActivity) > (time.Second * 60) {
			delete(sessions, k)
		}
	}
	fmt.Println("After Clean")
	showSessions()
}
func showSessions(){
	fmt.Println("************")
	for k,v := range sessions {
		fmt.Println(k, v.username)
	}
	fmt.Println("")
}
type session struct {
	username string
	lastActivity time.Time
}
type user struct {
	Username string
	FirstName string
	LastName string
	Password []byte
}