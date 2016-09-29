package main

import (
	"net/http"
	"html/template"
	"encoding/json"
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// declare page type
type Page struct {
	Title     string
	ChartTitle string
	ChartData string
}

// declare chart serie type
type Serie struct {
	Name  string    `json:"name"`
	Value int    `json:"value"`
}

// define db and page template
var db *sql.DB
var pageTemplate *template.Template

// init db and page template
func init() {
	pageTemplate = template.Must(template.ParseFiles("templates/index.html"))
	var err error
	db, err = sql.Open("mysql", "anychart_user:anychart_pass@/anychart_db")
	if err != nil {
		panic(err)
	}
}

// get top fruits from database
func getFruits(count int) []Serie {
	res, err := db.Query("SELECT * FROM fruits ORDER BY value DESC LIMIT ?", count)
	if err != nil {
		panic(err)
	}
	var	(
		id, value int
		name string
		series []Serie
	)
	for res.Next() {
		err = res.Scan(&id, &name, &value)
		if err != nil {
			panic(err)
		}
		series = append(series, Serie{name, value})
	}
	return series
}

func handler(w http.ResponseWriter, r *http.Request) {
	// get chart data
	data := getFruits(5)
	// decode chart data to JSON
	b, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	// init page data and run template
	page := &Page{
		Title: "Anychart Golang example",
		ChartTitle: "Top 5 fruits",
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
	defer db.Close()
}

