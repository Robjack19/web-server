package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/index.html")
	fmt.Println(t.Execute(w, nil))
}

func info(w http.ResponseWriter, r *http.Request) {

	t, _ := template.ParseFiles("templates/info.html")
	fmt.Println(t.Execute(w, nil))
}

func robjack(w http.ResponseWriter, r *http.Request) {

	t, _ := template.ParseFiles("templates/robjack.html")
	fmt.Println(t.Execute(w, nil))
}

func main() {

	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./static"))))
	http.HandleFunc("/", index)
	http.HandleFunc("/info", info)
	http.HandleFunc("/robjack", robjack)
	http.ListenAndServe(":8080", nil)
}
