package main

import (
    "encoding/json"
    "log"
    "net/http"
)

type HealthResponse struct {
    Status string `json:"status"`
}

type Item struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/api/health", withCORS(healthHandler))
    mux.HandleFunc("/api/items", withCORS(itemsHandler))
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodOptions {
            addCORSHeaders(w)
            w.WriteHeader(http.StatusNoContent)
            return
        }
        http.NotFound(w, r)
    })

    addr := ":8080"
    log.Printf("Backend listening on %s", addr)
    log.Fatal(http.ListenAndServe(addr, mux))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
        return
    }
    addCORSHeaders(w)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(HealthResponse{Status: "ok"})
}

func itemsHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
        return
    }
    addCORSHeaders(w)
    w.Header().Set("Content-Type", "application/json")
    items := []Item{{1, "First item"}, {2, "Second item"}}
    json.NewEncoder(w).Encode(items)
}

func withCORS(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        addCORSHeaders(w)
        if r.Method == http.MethodOptions {
            w.WriteHeader(http.StatusNoContent)
            return
        }
        next(w, r)
    }
}

func addCORSHeaders(w http.ResponseWriter) {
    w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
    w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
}
