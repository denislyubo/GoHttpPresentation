package main

import (
	"net/http"
)

func handleFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"text":"welcome"}`))
}

func main() {
	http.HandleFunc("/", handleFunc)
	http.ListenAndServe(":8080", nil)
}
