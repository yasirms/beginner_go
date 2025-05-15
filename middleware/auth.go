package middleware
import (
	"fmt"
	"net/http"
	"time"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)
// Claims struct to hold JWT claims
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.JSON(401, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}
		// Here you would typically validate the token and extract user information

		if len(token) < 7 || authHeader[:7] != "Bearer" {
			 // Example validation
			c.JSON(401, gin.H{"error": "Invalid Authorization format"})
			c.Abort()
			return
	}
	tokenString := token[7:]

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("your-256-bit-secret"), nil
	}
	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid token"})
		c.Abort()
		return	
}
if !token.Valid {
		c.JSON(401, gin.H{"error": "Invalid token"})
		c.Abort()
		return
	}
	c.Set("claims", claims)
	c.Next()
}
}
