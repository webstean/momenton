package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// subsituted at compile time
var Lastcommitsha string

type Version struct {
	Version       string
	Lastcommitsha string
	Description   string
}

func webversion(w http.ResponseWriter, r *http.Request) {
	appversion := Version{"1.0", Lastcommitsha, "pre-interview technical test"}

	js, err := json.Marshal(appversion)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func main() {
	fmt.Printf("Starting Web Server...(%s)", Lastcommitsha)
	http.HandleFunc("/version", webversion)
	http.ListenAndServe(":3000", nil)
}
