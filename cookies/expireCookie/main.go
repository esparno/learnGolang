package main

import (
	"net/http"
	"fmt"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/set", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/expire", expire)
	http.ListenAndServe(":8080", nil)
}
func index(wr http.ResponseWriter, req *http.Request ){
	fmt.Fprint(wr, `<a href="/set">Set Cookie</a><br /><a href="/read">Read Cookie</a><br /><a href="/expire">Expire Cookie</a> `)
}
func set(wr http.ResponseWriter, req *http.Request) {
	cookie := &http.Cookie{ Name: "my-cookie", Value: "Erin"}
	http.SetCookie(wr, cookie)
	http.Redirect(wr, req, "/", http.StatusSeeOther)
}
func read(wr http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("my-cookie")
	if err != nil {
		fmt.Fprint(wr, `<div>No Cookies</div>`)
		fmt.Fprint(wr, `<a href="/">Index</a> `)
		return
	}
	fmt.Fprint(wr, `<a href="/">Index</a><br /> `)
	fmt.Fprintf(wr, "Cookie Name: %s Cookie Value: %s", cookie.Name, cookie.Value)
}
func expire(wr http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("my-cookie")
	if err != nil {
		http.Redirect(wr, req, "/set", http.StatusSeeOther)
		return
	}
	cookie.MaxAge = -1
	http.SetCookie(wr, cookie)
	fmt.Fprint(wr, `<a href="/">Index</a> `)
}