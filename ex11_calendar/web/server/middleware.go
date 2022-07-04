package server

import (
	"log"
	"net/http"
)

// Промежуточный метод для логирования событий
func middleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request: Method " + r.Method + ", Url " + r.RequestURI)
		next.ServeHTTP(w, r)
	}
}
