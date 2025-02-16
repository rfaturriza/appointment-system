package handlers

import (
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/rfaturriza/appointment-system/db"
	"github.com/rfaturriza/appointment-system/models"
	"gorm.io/gorm"
)

func Login(c *gin.Context) {
    username := c.PostForm("username")

    if username == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Username is required"})
        return
    }

    result := db.DB.First(&models.User{}, "username = ?", username)
    if result.Error != nil {
       if errors.Is(result.Error, gorm.ErrRecordNotFound){
         c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
         return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query database"})
    }
    expiresIn := os.Getenv("JWT_EXPIRES_IN")
    duration, err := time.ParseDuration(expiresIn)
    if err != nil {
        // Default to 24 hours if parsing fails
        duration = 24 * time.Hour
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "username": username,
        "exp":      time.Now().Add(duration).Unix(),
    })

    jwtSecret := os.Getenv("JWT_SECRET")

    tokenString, err := token.SignedString([]byte(jwtSecret))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": tokenString})
}