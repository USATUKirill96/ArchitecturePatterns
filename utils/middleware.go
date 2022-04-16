package utils

import (
	"fmt"
	"net/http"
)

func RecoverPanic(logger *Logger) func(http.Handler) http.Handler {
	recoveryFunc := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					logger.ServerError(fmt.Errorf("%s", err), w)
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
	return recoveryFunc
}
