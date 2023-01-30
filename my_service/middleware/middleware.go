package middleware

import (
	"RestAPI/pkg/logger"
	"net/http"
	"time"
)

var l = logger.Get()

func Midd(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()

		l.
			Info().
			Str("method", r.Method).
			Str("url", r.URL.RequestURI()).
			Str("user_agent", r.UserAgent()).
			Dur("elapsed_ms", time.Since(start)).
			Msg("incoming request")

		next.ServeHTTP(w, r)
	}
}
