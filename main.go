package main

import (
	"net/http"
	"html/template"
	"fmt"
)

// template variable is accessed from multiple files. So it's better to create a global variable
var tmpl *template.Template

func main(){
	// Define Address (IP:Port)
	addr := ":8080"

	// Build the Router
	mux := http.NewServeMux()

	// Set Routes
	mux.HandleFunc("/", catchAll)
	
	// Parse the templates in advance and store it in a Global Variable
	tmpl = template.Must(template.ParseGlob("templates/*.html"))

	
	// Start the listener
	fmt.Printf("Listening on %s...\n", addr)
	http.ListenAndServe(addr, mux)
}
