package main

import (
	"fmt"
	"net/http"
)

type OurHandler struct{}

func (h OurHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `{"text":"welcome"}`)
}

func main() {
	http.ListenAndServe(":8080", OurHandler{})
}
