package middlewares

import (
	_ "go-phonebooks/models"
	"net/http"
	// u "go-phonebooks/utils"
	// "strings"
	// jwt "github.com/dgrijalva/jwt-go"
	// "os"
	// "context"
	// "fmt"
)

var JwtAuthMiddleware = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// noAuth := []string{"/api/user/new", "/api/user/login"}
		// requestPath := r.URL.Path
		next.ServeHTTP(w, r)
		// for _, value := range noAuth {
		// 	if value == requestPath {
		// 		next.ServeHttp(w, r)
		// 		return
		// 	}
		// }
	})
}
