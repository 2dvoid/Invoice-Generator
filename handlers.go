package main

import (
	"net/http"
)

func catchAll(w http.ResponseWriter, r *http.Request){
	tmpl.ExecuteTemplate(w, "home.tmpl.html", nil)
}
