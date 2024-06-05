//midlwares/logging.go
package middlewares

import (
    "net/http"
    "log"
)

func Logging(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Do stuff before
        log.Println("Before request middleware")

        // Call the next handler, which can be another middleware in the chain, or the final handler.
        next.ServeHTTP(w, r)

        // Do stuff after
        log.Println("After request middleware")
    })
}