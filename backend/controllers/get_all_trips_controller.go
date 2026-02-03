package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)


func (tc *TripController) GetAllTrips(context *gin.Context) {
	trips, err := tc.repo.GetAllTrips()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch events"})
		log.Printf("error: %v", err)
		return
	}
	context.JSON(http.StatusOK, trips)
}

