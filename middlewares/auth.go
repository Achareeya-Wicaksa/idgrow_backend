// middleware/auth.go
package middleware

import (
    "net/http"
    "strings"

    "github.com/dgrijalva/jwt-go"
)

func JWTAuth(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        authHeader := r.Header.Get("Authorization")
        if authHeader == "" {
            http.Error(w, "Authorization header is missing", http.StatusUnauthorized)
            return
        }

        tokenString := strings.Split(authHeader, "Bearer ")[1]
        if tokenString == "" {
            http.Error(w, "Token is missing", http.StatusUnauthorized)
            return
        }

        // Verifikasi token JWT
        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            return []byte("secret_key"), nil // Ganti dengan kunci yang Anda gunakan untuk membuat token
        })

        if err != nil || !token.Valid {
            http.Error(w, "Invalid token", http.StatusUnauthorized)
            return
        }

        next.ServeHTTP(w, r)
    })
}
