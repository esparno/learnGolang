package main

import (
	"net/http"
	"github.com/satori/go.uuid"
	"fmt"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ido", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
func foo (wr http.ResponseWriter, req *http.Request){
	cookie, err := req.Cookie("session")
	if err != nil {
		cookie = &http.Cookie{
			Name: "session",
			Value: uuid.NewV4().String(),
			HttpOnly: true, // makes it so cookie is not accessible using javascript
		}
		http.SetCookie(wr, cookie)
	}
	fmt.Println(cookie)
}