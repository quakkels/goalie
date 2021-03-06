package controllers

import (
	"net/http"
)

// HelloController handles the ubiquitus "hello world" greeting
func HelloController(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	w.Write([]byte("Hello, World!"))
}

// HelloProtected is here
func HelloProtected(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	w.Write([]byte("You are authenticated."))
}
