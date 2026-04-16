package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/privacy", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	log.Println("Ran http://localhost:8082/privacy")

	err := http.ListenAndServe(":8082", nil)
	if err != nil {
		log.Fatal("error: ", err)
	}
}
