package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"time"
)


func handleShorten(w http.ResponseWriter, r *http.Request) {
	longURL := r.FormValue("url")
	shortURL, slug := generateShortURL(r)

	saveURL(slug, longURL)

	renderTemplate(w, "index.html", map[string]string{
		"ShortURL": shortURL,
		"LongURL": longURL,
	})
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	tmplPath := fmt.Sprintf("templates/%s", tmpl)

	// Parse HTML
	t, err := template.ParseFiles(tmplPath)

	if err != nil {
		http.Error(w, "Error Parsing HTML", http.StatusInternalServerError)
		return
	}

	// Render HTML
	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, "Error Rendering Template", http.StatusInternalServerError)
		return
	}
}

func generateShortURL(r *http.Request) (string, string) {
	rand.NewSource(time.Now().UnixNano())
	domain := r.Host
	randomString := generateRandomString(6)
	shortURL := fmt.Sprintf("http://%s/%s", domain, randomString)
	return shortURL, randomString
}

func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}