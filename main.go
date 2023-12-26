package main

import (
	"database/sql"
	"fmt"
	"net/http"

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
	router.HandleFunc("/", indexHandler)
	router.HandleFunc("/{short}", redirectHandler)
	http.Handle("/", router)

	// Handle all assets in the assets folder
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))

	// Start the server
	http.ListenAndServe(":8080", nil)
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

