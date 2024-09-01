// middlewares/auth.go
package middlewares

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "github.com/dgrijalva/jwt-go"
    "backend/controllers"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString := c.GetHeader("Authorization")

        if tokenString == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Request does not contain an access token"})
            c.Abort()
            return
        }

        claims := &controllers.Claims{}
        token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
            return []byte("secret_key"), nil
        })

        if err != nil || !token.Valid {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        c.Next()
    }
}
