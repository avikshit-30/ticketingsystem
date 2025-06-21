package middleware

import (
	"log"
	"net/http"
	"time"
)

func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		log.Printf("📥 %s %s from %s", r.Method, r.URL.Path, r.RemoteAddr)

		next.ServeHTTP(w, r)

		log.Printf("✅ %s %s completed in %v\n", r.Method, r.URL.Path, time.Since(start))
	})
}
