package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {

	serverMux := http.NewServeMux()

	serverMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("home page")
		t, err := template.ParseFiles("templates/home.html")
		if err != nil {
			log.Fatalln(err.Error())
		}
		t.Execute(w, nil)
	})

	serverMux.HandleFunc("/initiate-payment", func(w http.ResponseWriter, r *http.Request) {
		log.Println("initiate payment")
		t, err := template.ParseFiles("templates/initiate-payment.html")
		if err != nil {
			log.Fatalln(err.Error())
		}
		t.Execute(w, nil)
	})

	serverMux.HandleFunc("/acknowledge-payment", func(w http.ResponseWriter, r *http.Request) {
		log.Println("acknowledge-payment")
	})

	http.ListenAndServe(":9999", serverMux)
}
