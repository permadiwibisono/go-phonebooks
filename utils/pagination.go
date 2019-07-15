package utils

import "net/http"

type PaginationQuery struct {
	QueryParams
	Page    int
	PerPage int
}

type QueryParams struct {
	Queries map[string][]string
}

func GetPaginationQueryParams(r *http.Request) PaginationQuery {
	getPaginationQuery := r.Context().Value("pagination")
	pagination := PaginationQuery{
		Page:    1,
		PerPage: 16,
	}
	if getPaginationQuery != nil {
		pagination = getPaginationQuery.(PaginationQuery)
	}
	return pagination
}
