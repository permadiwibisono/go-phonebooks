package middlewares

import (
	"context"
	"fmt"
	u "go-phonebooks/utils"
	"net/http"
	"strconv"
	"time"
)

var PaginationQueryParams = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			queries := r.URL.Query()
			if len(queries) > 0 {
				// var perPage int64 = 16
				// var page int64 = 1
				pagination := u.PaginationQuery{
					Page:    1,
					PerPage: 16,
				}
				modQueries := make(map[string][]string)

				for key, item := range queries {
					fmt.Printf("%s %s", key, item)
					modQueries[key] = item
				}
				if i, ok := modQueries["page"]; ok {
					p, err := strconv.ParseInt(i[0], 10, 64)
					fmt.Printf("%d \n", p)
					if err == nil {
						pagination.Page = int(p)
					}
				}
				if i, ok := modQueries["per_page"]; ok {
					p, err := strconv.ParseInt(i[0], 10, 64)
					fmt.Printf("%d \n", p)
					if err == nil {
						pagination.PerPage = int(p)
					}
				}
				pagination.QueryParams.Queries = modQueries
				ctx := context.WithValue(r.Context(), "pagination", pagination)
				r = r.WithContext(ctx)
				fmt.Printf("[%s] %s %s\n", time.Now().Format(time.RFC822Z), "Your query params", queries)
			}
		}
		next.ServeHTTP(w, r)
	})
}
