package main

import (
	"database/sql"
	"fmt"
	"sync"
)

type URLInfo struct {
	LongURL string
}

// create a mutex
var urlsLock sync.RWMutex

// Save the newly generated short URLs and original URLS into the database
func saveURL(short, long string) {
	urlsLock.Lock()
	defer urlsLock.Unlock()

	_, err := db.Exec("INSERT INTO urls (short_url, long_url) VALUES (?, ?)", short, long)
	if err != nil {
		fmt.Println("Error saving URL:", err)
	}
}


func getURLInfo(short string) (URLInfo, bool) {
	urlsLock.RLock()
	defer urlsLock.RUnlock()

	var info URLInfo
	err := db.QueryRow("SELECT long_url FROM urls WHERE short_url = ?", short).Scan(&info.LongURL)
	if err == sql.ErrNoRows {
		return info, false
	} else if err != nil {
		fmt.Println("Error retrieving URL info:", err)
	}

	return info, true
}