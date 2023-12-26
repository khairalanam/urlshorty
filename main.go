package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func main() {
	// Open SQLite database
	var err error
	fmt.Println("Hello")
	db, err = sql.Open("sqlite3", "url-shortener.db")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}

	// Create a new router and endpoints
	router := mux.NewRouter()
	router.Use(CORSMiddleware)
	router.HandleFunc("/", indexHandler)
	router.HandleFunc("/{short}", redirectHandler)
	http.Handle("/", router)

	// Handle all assets in the assets folder
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))

	// for Railway deployment
	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	// Start the server
	http.ListenAndServe("0.0.0.0:" + port, nil)
}

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		handleShorten(w, r)
		return
	}

	renderTemplate(w, "index.html", nil)
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	// Slug extraction
    vars := mux.Vars(r)
    slug := vars["short"]

	// Fetch the actual URL
    info, found := getURLInfo(slug)

    if found {
        fmt.Println("Redirecting to:", info.LongURL)
        http.Redirect(w, r, info.LongURL, http.StatusFound)
    } else {
        fmt.Println("URL not found, returning 404")
        http.NotFound(w, r)
    }
}

