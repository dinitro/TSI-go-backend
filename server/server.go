package server

import (
	"log"
	"net/http"
	"os"
)

func Init() {
	var port string
	if port = os.Getenv("PORT"); port == "" {
		port = "8080"
	}

	router := Router()

	log.Printf("Server starting at http://localhost:%s\n", port)

	// Add CORS middleware
	handler := addCorsHeader(router)

	err := http.ListenAndServe(":"+port, handler)

	log.Fatal(err)
}

func addCorsHeader(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Allow all origins
		w.Header().Set("Access-Control-Allow-Origin", "*")

		// Allow specific headers
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Pass the request to the next handler
		handler.ServeHTTP(w, r)
	})
}
