package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)
	http.HandleFunc("/greet", greeter)
	http.HandleFunc("/time", timeHandler)

	log.Println("Listening...")
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func root(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("public/index.html")
	t.Execute(w, nil)
}

func greeter(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	t, _ := template.ParseFiles("public/greet.html")
	err := t.Execute(w, username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello. The time is: "+time.Now().Format(time.RFC850))
	fmt.Fprintln(w, r.UserAgent())
	fmt.Fprintln(w, r.URL.Path)
}
