package main

import (
	"log"
	"net/http"
	"os"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/home.html")
}

func aboutPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/about.html")
}

func contactPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/contact.html")
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	// Redirect to home
	http.Redirect(w, r, "/home", http.StatusFound)
}

func main() {
	// Set routes
	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/home", homePage)
	http.HandleFunc("/about", aboutPage)
	http.HandleFunc("/contact", contactPage)

	// Use environment variable for port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not set
	}

	log.Printf("Server starting on port %s...", port)
	err := http.ListenAndServe("0.0.0.0:"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
