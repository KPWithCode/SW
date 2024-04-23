package main

import (
	"net/http"
	"fmt"
)

func main() {
	mux := http.NewServeMux()

	// CORS middleware
	corsHandler := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
			w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			if r.Method == http.MethodOptions {
				// Handle preflight request
				w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
				w.WriteHeader(http.StatusOK)
				return
			}
			next.ServeHTTP(w, r)
		})
	}

	// CORS middleware
	mux.Handle("/", corsHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})))

	// Routes
	// routes.SOPRoutes(mux)

	fmt.Println("Server is now listening on :8080")

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
	// if err := server.ListenAndServeTLS("cert.pem", "key.pem"); err != nil {
	// 	panic(err)
	// }
}
