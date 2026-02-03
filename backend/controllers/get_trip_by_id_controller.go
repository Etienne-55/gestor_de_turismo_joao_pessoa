package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)


func (tc *TripController) GetTripById(context *gin.Context) {
	tripID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse the trip"})
		log.Printf("error: %v", err)
		return
	} 

	trip, err := tc.repo.GetTripById(tripID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch the trip info"})
		log.Printf("error: %v", err)
		return
	}
	context.JSON(http.StatusOK, trip)
}

