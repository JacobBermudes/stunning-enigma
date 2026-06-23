package main

import (
	"html/template"
	"log"
	"net/http"
)

type SuccessData struct {
	Days string
	Tier string
}

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

	http.HandleFunc("/terms", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "terms.html")
	})

	http.HandleFunc("/success", func(w http.ResponseWriter, r *http.Request) {
		days := r.URL.Query().Get("days")
		tier := r.URL.Query().Get("tier")

		if days == "" {
			days = "30"
		}
		if tier == "" {
			tier = "Basic"
		}

		data := SuccessData{
			Days: days,
			Tier: tier,
		}

		tmpl, err := template.ParseFiles("success.html")
		if err != nil {
			http.Error(w, "Ошибка сервера при загрузке страницы", http.StatusInternalServerError)
			log.Println("Ошибка шаблона:", err)
			return
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			log.Println("Ошибка рендера шаблона:", err)
		}
	})

	http.Handle("/favicons/", http.StripPrefix("/favicons/", http.FileServer(http.Dir("favicons"))))

	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "favicons/favicon.ico")
    })
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.Redirect(w, r, "/about", http.StatusFound)
			return
		}
		http.Redirect(w, r, "/about", http.StatusFound)
	})

	log.Println("Ran http://localhost:8082")

	err := http.ListenAndServe(":8082", nil)
	if err != nil {
		log.Fatal("error: ", err)
	}
}
