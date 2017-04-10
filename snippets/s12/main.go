package main

import (
	"encoding/json"
	"net/http"
)

// Workshop defines an event.
type Workshop struct {
	Title        string `json:"title"`
	Participants int    `json:"count"`
}

func main() {
	lgworkshop := Workshop{
		Title:        "The Let's Go workshop",
		Participants: 33,
	}
	http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(lgworkshop)
	})
	http.ListenAndServe(":9876", nil)
}
