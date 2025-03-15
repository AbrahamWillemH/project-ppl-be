package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// AuthMiddleware checks JWT token validity and extracts claims
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			c.Abort()
			return
		}

		// Remove "Bearer " prefix if present
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		secretKey := os.Getenv("JWT_SECRET")
		if secretKey == "" {
			secretKey = "defaultsecret"
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Extract claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		// Set user role in context for further middleware
		if role, exists := claims["role"].(string); exists {
			c.Set("role", role)
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Role not found in token"})
			c.Abort()
			return
		}

		c.Next()
	}
}

// AdminMiddleware allows only admin role
func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists || role.(string) != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden: Admins only"})
			c.Abort()
			return
		}
		c.Next()
	}
}

// StudentMiddleware allows students and admins
func StudentMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists || (role.(string) != "student" && role.(string) != "admin") {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden: Students and Admins only"})
			c.Abort()
			return
		}
		c.Next()
	}
}

// TeacherMiddleware allows teachers and admins
func TeacherMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists || (role.(string) != "teacher" && role.(string) != "admin") {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden: Teachers and Admins only"})
			c.Abort()
			return
		}
		c.Next()
	}
}
