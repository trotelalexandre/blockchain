package handlers

import (
	"net/http"
)

func Handler(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173") // Allow requests from your Vite app
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST")   		   // Allow necessary methods
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")         // Allow necessary headers

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next(w, r)
	}
}
