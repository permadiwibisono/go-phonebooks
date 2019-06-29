package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type PaginationQuery struct {
	Page    int64
	PerPage int64
}

var PaginationQueryParams = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			queries := r.URL.Query()
			if len(queries) > 0 {
				var perPage int64 = 16
				var page int64 = 1
				modQueries := make(map[string][]string)

				for key, item := range queries {
					fmt.Printf("%s %s", key, item)
					if key != "per_page" || key != "page" {
						modQueries[key] = item
					}
				}
				if i, ok := queries["page"]; ok {
					p, err := strconv.ParseInt(i[0], 10, 64)
					if err == nil {
						perPage = p
					}
				}
				if i, ok := queries["per_page"]; ok {
					p, err := strconv.ParseInt(i[0], 10, 64)
					if err == nil {
						perPage = p
					}
				}
				ctx := context.WithValue(r.Context(), "pagination", PaginationQuery{Page: page, PerPage: perPage})
				r = r.WithContext(ctx)
				fmt.Printf("[%s] %s %s\n", time.Now().Format(time.RFC822Z), "Your query params", queries)
			}

		}
		next.ServeHTTP(w, r)
	})
}
