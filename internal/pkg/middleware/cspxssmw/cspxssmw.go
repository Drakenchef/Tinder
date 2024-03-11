package cspxssmw

import (
	"html"
	"net/http"
)

type CspXssMW struct {
}

func NewCspXssMW() *CspXssMW {
	return &CspXssMW{}
}

func (cx *CspXssMW) MiddlewareCSP(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		csp := "default-src 'self'; script-src 'self'; style-src 'self'; form-action 'self'; base-uri 'self'; plugin-types 'none';"
		w.Header().Set("Content-Security-Policy", csp)
		next.ServeHTTP(w, r)
	})
}

func (cx *CspXssMW) MiddlewareXSS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sanitizeQueryParams(r)
		next.ServeHTTP(w, r)
	})
}

func sanitizeQueryParams(r *http.Request) {
	query := r.URL.Query()

	for key := range query {
		for i, val := range query[key] {
			query.Set(key, html.EscapeString(val))
			query[key][i] = html.EscapeString(val)
		}
	}

	r.URL.RawQuery = query.Encode()
}
