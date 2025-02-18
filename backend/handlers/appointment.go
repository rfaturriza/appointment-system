package handlers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofrs/uuid"
	"github.com/rfaturriza/appointment-system/db"
	"github.com/rfaturriza/appointment-system/models"

	"github.com/gin-gonic/gin"
)
func CreateAppointment(c *gin.Context) {
    if db.DB == nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "database connection not initialized"})
        return
    }
    var appointment models.Appointment
    if err := c.ShouldBindJSON(&appointment); err != nil {
        log.Println(err)
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid input",
        })
        return
    }
    var user models.User
    claims, _ := c.Get("claims")
    username := claims.(jwt.MapClaims)["username"].(string)
    result := db.DB.First(&user, "username = ?", username)
    if result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }
    timezone, err := time.LoadLocation(user.PreferredTimezone)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid timezone"})
        return
    }

    // convert appointment start and end time to user's preferred timezone
    appointment.Start = appointment.Start.In(timezone)
    appointment.End = appointment.End.In(timezone)

    // Validate working hours and timezone
    isValidAppointmentTime := isValidAppointmentTime(appointment.Start, appointment.End)
    if !isValidAppointmentTime {
        c.JSON(http.StatusBadRequest, gin.H{"error": "start & end time must be within working hours (8AM - 5PM) and start time must be before end time"})
        return
    }
    appointment.ID = uuid.Must(uuid.NewV4())
    appointment.CreatorID = user.ID

    // save appointment start and end time in UTC
    appointment.Start = appointment.Start.UTC()
    appointment.End = appointment.End.UTC()
    db.DB.Create(&appointment)
    c.JSON(http.StatusCreated, appointment)
}

func isValidAppointmentTime(start, end time.Time) bool {
    if start.After(end) {
        return false
    }
    if start.Hour() < 8 || start.Hour() >= 17 {
        return false
    }

    if end.Hour() < 8 || end.Hour() >= 17 {
        return false
    }
    return true
}

func GetUpcomingAppointments(c *gin.Context) {
    var user models.User
    claims, _ := c.Get("claims")
    username := claims.(jwt.MapClaims)["username"].(string)
    resultUser := db.DB.First(&user, "username = ?", username)
    if resultUser.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }
    timezone, err := time.LoadLocation(user.PreferredTimezone)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid timezone"})
        return
    }
    var appointments []models.Appointment
    timeNowTz:= time.Now().In(timezone)
    timeNowInUTC := timeNowTz.UTC()
    fmt.Println(timeNowTz)
    result := db.DB.Preload("Creator").Where("start > ?", timeNowInUTC).Order("start ASC").Find(&appointments)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching appointments"})
        return
    }
    c.JSON(http.StatusOK, appointments)
}