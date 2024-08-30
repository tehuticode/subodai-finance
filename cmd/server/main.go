package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/tehuticode/subodai-finance/internal/auth"
	"github.com/tehuticode/subodai-finance/internal/database"
	"github.com/tehuticode/subodai-finance/internal/trading"
)

func main() {
	err := database.InitDB()
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	// Serve static files
	fs := http.FileServer(http.Dir("web/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Home page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("web/templates/layout.html", "web/templates/index.html"))
		tmpl.Execute(w, nil)
	})

	// Trade page
	http.HandleFunc("/trade", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("web/templates/layout.html", "web/templates/trade.html"))
		tmpl.Execute(w, nil)
	})

	// Authentication routes
	http.HandleFunc("/register", auth.RegisterHandler)
	// TODO: Add login route

	// API routes
	http.HandleFunc("/api/market-summary", trading.MarketSummaryHandler)
	http.HandleFunc("/api/place-trade", trading.PlaceTradeHandler)

	log.Println("Starting server on :8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
