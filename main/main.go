package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func MainHandler(resp http.ResponseWriter, r *http.Request) {
	io.WriteString(resp, time.Now().Format("2006-01-02 15:04:05"))
}

func AboutMe(resp http.ResponseWriter, r *http.Request) {
	fmt.Println("I am Maxime!")
}

func main() {
	http.HandleFunc("/", MainHandler)
	http.HandleFunc("/about", AboutMe)
	fmt.Println("Listening on port 5050...")
	http.ListenAndServe(":5050", nil)
}
