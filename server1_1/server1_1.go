package main

import (
	"fmt"
	"net/http"
)

func handleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `{"text":"welcome"}`)
}

func main() {
	http.HandleFunc("/", handleFunc)
	http.ListenAndServe(":8080", nil)
}
