package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// static files
	mux.Handle("GET /", http.FileServer(http.Dir("web")))

	// /room/JQ7B
	mux.HandleFunc("GET /room/{code}", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/index.html")
	})

	log.Println("listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
