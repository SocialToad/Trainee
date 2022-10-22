package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/test1", server) // заметка.тест
	http.HandleFunc("/", server1)     // заметка.текст
	http.ListenAndServe(":8888", nil)
}
func server(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "readtext")
}
func server1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "smiletext")
}
