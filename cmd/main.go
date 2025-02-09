package main

import (
	"fmt"
	"blog-service-go/pkg/handler"
	"log"
	"net/http"
)

func main() {
	// Serve static files like HTML, CSS, and JS
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	// Set up the /users route
	http.HandleFunc("/users", handler.UsersHandler)

	// Set up the root route for the HTML page
	http.HandleFunc("/", handler.IndexHandler)

	fmt.Println("Server is running on port http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil)) // Start the server
}
