package main

import (
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
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Welcome to Subodai Finance!"))
    })
    http.HandleFunc("/register", auth.RegisterHandler)
    log.Println("Starting server on :8080")
    err = http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatalf("Error starting server: %v", err)
    }
}
