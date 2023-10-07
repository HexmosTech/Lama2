package main

import (
	"log"
	"net/http"
	"time"
)

func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s %v", r.RemoteAddr, r.Method, r.URL, time.Since(startTime))
	})
}

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", logger(fs))

	log.Println("Listening on http://localhost:3000/index.html")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
