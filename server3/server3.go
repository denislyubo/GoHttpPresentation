package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type GreetController struct{}

func (gc *GreetController) GreetStranger(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"text":"welcome stranger"}`))
}

func (gc *GreetController) GreetPerson(w http.ResponseWriter, r *http.Request) {
	var name string
	if r.ParseForm() != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"text":"unable to parse form"}`))
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
}

func (gc *GreetController) GreetPersonURL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	str := vars["name"]
	w.Write([]byte(`{"text":"welcome ` + str + `"}`))
}

func main() {
	r := mux.NewRouter()

	greetC := GreetController{}

	r.HandleFunc("/greeting", greetC.GreetStranger).Methods("GET")
	r.HandleFunc("/greeting", greetC.GreetPerson).Methods("POST")
	r.HandleFunc("/greeting/{name:[a-z]+}", greetC.GreetPersonURL).Methods("POST")

	http.ListenAndServe(":8080", r)
}
