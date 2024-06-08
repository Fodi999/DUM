//logging.go
package middlewares

import (
    "log"
    "net/http"
)

func Logging(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Println("Before request middleware")
        next.ServeHTTP(w, r)
        log.Println("After request middleware")
    })
}




