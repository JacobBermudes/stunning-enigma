package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/logo.webp", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "logo.webp")
	})

	http.HandleFunc("/privacy", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "privacy.html")
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "about.html")
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/about", http.StatusFound)
	})

	log.Println("Ran http://localhost:8082")

	err := http.ListenAndServe(":8082", nil)
	if err != nil {
		log.Fatal("error: ", err)
	}
}
