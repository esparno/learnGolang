package main

import (
	"net/http"
	"strconv"
	"log"
	"fmt"
	"io"
)

func main() {
	http.HandleFunc("/", countCookie)
	http.ListenAndServe(":8080", nil)
}
func countCookie(wr http.ResponseWriter, req *http.Request){
	c, err := req.Cookie("count")
	cookie := &http.Cookie{}
	if err == http.ErrNoCookie {
		cookie.Name = "count"
		cookie.Value = "0"
	} else {
		count, err := strconv.Atoi(c.Value)
		if err != nil {
			log.Fatalln(err)
		}
		count ++
		cookie.Value = strconv.Itoa(count)
		cookie.Name = "count"
	}
	http.SetCookie(wr, cookie)
	fmt.Println(cookie.Value)
	io.WriteString(wr, cookie.Value)


}