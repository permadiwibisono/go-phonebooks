package middlewares

import (
	"fmt"
	"net/http"
	"time"
)

var Logger = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("[%s] %s %s\n", time.Now().Format(time.RFC822Z), r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
