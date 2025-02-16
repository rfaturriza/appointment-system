package handlers

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/rfaturriza/appointment-system/db"
	"github.com/rfaturriza/appointment-system/models"
)

func GetUser(c *gin.Context) {
    userID := c.Param("id")
    var user models.User

    result := db.DB.Where("id = ?", userID).First(&user)
    if result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    c.JSON(http.StatusOK, user)
}

func GetCurrentUser(c *gin.Context) {
    claims := c.MustGet("claims").(jwt.MapClaims)
    username := claims["username"].(string)
    var user models.User

    result := db.DB.Where("username = ?", username).First(&user)
    if result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    c.JSON(http.StatusOK, user)
}