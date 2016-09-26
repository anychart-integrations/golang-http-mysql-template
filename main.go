package main

import (
	"net/http"
	"html/template"
	"encoding/json"
	"log"
)

// declare page type
type Page struct {
	Title string
	ChartData string
}

// declare chart serie type
type Serie struct {
	Name  string    `json:"name"`
	Value int    `json:"value"`
}

// define and init page template
var pageTemplate *template.Template

func init() {
	pageTemplate = template.Must(template.ParseFiles("templates/index.html"))
}

func handler(w http.ResponseWriter, r *http.Request) {
	// assuming your have data
	data := []Serie{
		{"Chocolate", 5},
		{"Rhubarb compote", 2},
		{"Crêpe Suzette", 2},
		{"American blueberry", 2},
		{"Buttermilk", 1},
	}
	// decode chart data to JSON
	b, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	// init page data and run template
	page := &Page{
		Title: "Anychart Golang example",
		ChartData: string(b),
	}
	pageTemplate.Execute(w, page)
}


func main() {
	// define routes
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/static/", fs)
	http.HandleFunc("/", handler)
	// start web server
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

