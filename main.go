package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

var apiKey = ""

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error reading .env file")
	}
	apiKey = os.Getenv("POLYGON_KEY")

	http.Handle("/static/",
		http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("./templates/index.html"))
		tmpl.Execute(w, nil)
	})

	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("./templates/fragments/results.html"))
		data := map[string][]Stock{
			"Results": SearchTicker(r.URL.Query().Get("key"), apiKey),
		}
		tmpl.Execute(w, data)
	})

	http.HandleFunc("/stock/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			ticker := r.PostFormValue("ticker")
			stk := SearchTicker(ticker, apiKey)[0]
			val := GetDailyValues(ticker, apiKey)
			tmpl := template.Must(template.ParseFiles("./templates/index.html"))
			tmpl.ExecuteTemplate(w, "stock-element",
				Stock{Ticker: stk.Ticker, Name: stk.Name, Price: val.Open})
		}
	})

	log.Println("App running on 8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
