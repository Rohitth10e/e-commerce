package middleware

import (
    "net/http"
    "strings"

    "github/rohitth10e/utils"
    "github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        // Get the Authorization header
        authHeader := ctx.GetHeader("Authorization")
        if authHeader == "" {
            ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
            return
        }

        // Expect "Bearer <token>"
        parts := strings.Split(authHeader, " ")
        if len(parts) != 2 || parts[0] != "Bearer" {
            ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
            return
        }

        tokenString := parts[1]

        // Verify token
        claims, err := utils.Verify(tokenString)
        if err != nil {
            ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            return
        }

        // Attach claims to context
        ctx.Set("email", claims.Email)
        ctx.Set("role", claims.Role)

        // Continue to next handler
        ctx.Next()
    }
}