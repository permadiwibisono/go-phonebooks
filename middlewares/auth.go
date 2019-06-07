package middlewares

import (
	"go-phonebooks/models"
	_ "go-phonebooks/models"
	"net/http"
	"os"

	u "go-phonebooks/utils"
	"strings"

	"context"
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
)

var JwtAuthMiddleware = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// noAuth := []string{"/api/user/new", "/api/user/login"}
		// requestPath := r.URL.Path
		fmt.Println("Auth jwt middleware entered!!")
		getTokenHeader := r.Header.Get("Authorization")
		if getTokenHeader == "" {
			u.RespondError(w, http.StatusForbidden, "Missing auth token!", nil)
			return
		}
		tokenSplitted := strings.Split(getTokenHeader, " ")
		if len(tokenSplitted) != 2 {
			u.RespondError(w, http.StatusBadRequest, "Invalid/Malformed auth token!", nil)
			return
		}
		tokenPart := tokenSplitted[len(tokenSplitted)-1]
		tk := &models.Token{}
		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("jwt_token")), nil
		})

		if err != nil {
			u.RespondError(w, http.StatusBadRequest, "Invalid/Malformed auth token!", nil)
			return
		}

		if !token.Valid {
			u.RespondError(w, http.StatusForbidden, "Invalid token!", nil)
			return
		}

		ctx := context.WithValue(r.Context(), "user", tk.UserID)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
		// for _, value := range noAuth {
		// 	if value == requestPath {
		// 		next.ServeHttp(w, r)
		// 		return
		// 	}
		// }
	})
}
var HelloMiddleware = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hello middleware entered!!")
		next.ServeHTTP(w, r)
	})
}
