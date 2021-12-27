package main

import (
	"fmt"
	"net/http"
)

func handleFunc(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"text":"Welcome stranger"}`))
	case "POST":
		var name string
		if r.ParseForm() != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"text":"unable to parse form"}`))
			return
		} else {
			rec, ok := r.Form["name"]
			if !ok {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(`{"text":"unable to parse form"}`))
				return
			}
			name = rec[0]
		}
		var resp []byte
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		resp = append(resp, []byte(fmt.Sprintf(`{"text":"Welcome %s"}`, name))...)
		w.Write(resp)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(`{"text":"method is not allowed"}`))
	}
}

func main() {
	http.HandleFunc("/greeting", handleFunc)
	http.ListenAndServe(":8080", nil)
}
