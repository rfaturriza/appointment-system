package middleware

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
	 tokenHeader := c.GetHeader("Authorization")
	 jwtSecret := os.Getenv("JWT_SECRET")
	 if tokenHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	 }

	 tokenSplit := strings.Split(tokenHeader, " ")
	 if len(tokenSplit) != 2 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	 }

	 tokenString := tokenSplit[1]
	 
	 token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, http.ErrAbortHandler
		}
			return []byte(jwtSecret), nil
	 })

	//  handle expired token
	if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorExpired != 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token Expired"})
			c.Abort()
			return
		}
	}
   
	 if err != nil || !token.Valid {
		log.Println(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	 }
   
	 if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
	  	c.Set("claims", claims)
	 } else {
	 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
	  	return
	 }
   
	 	c.Next()
	}
}