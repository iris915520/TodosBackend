package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var todos []string

func main() {
	http.HandleFunc("/list", func(w http.ResponseWriter, r *http.Request) {
		writeJson(w, todos)
	})
	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		var text = r.FormValue("text")
		todos = append(todos, text)
		http.Redirect(w, r, "/list", http.StatusSeeOther)
	})
	http.HandleFunc("/delete", func(w http.ResponseWriter, r *http.Request) {
		var text = r.FormValue("text")
		var newList = []string{}
		for _, todo := range todos {
			if todo != text {
				newList = append(newList, todo)
			}
		}
		todos = newList
		http.Redirect(w, r, "/list", http.StatusSeeOther)
	})
	http.HandleFunc("/update", func(w http.ResponseWriter, r *http.Request) {
		var oldText = r.FormValue("oldText")
		var newText = r.FormValue("newText")
		for i, todo := range todos {
			if todo == oldText {
				todos[i] = newText
			}
		}
		http.Redirect(w, r, "/list", http.StatusSeeOther)
	})
	fmt.Println("server started at http://localhost:80")
	e := http.ListenAndServe(":80", nil)
	if e != nil {
		fmt.Println(e)
	}
}

func writeJson(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	b, _ := json.MarshalIndent(data, "", "\t")
	w.Write(b)
}
