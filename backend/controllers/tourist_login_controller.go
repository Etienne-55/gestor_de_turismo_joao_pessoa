package controllers

import (
	"log"
	"net/http"
	"projeto_turismo_jp/models"
	"projeto_turismo_jp/utils"

	"github.com/gin-gonic/gin"
)


func (tc *TouristController) Login(context *gin.Context) {
	var tourist models.Tourist

	err := context.ShouldBindJSON(&tourist)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "invalid credentials 1"})
		return	
	}

	err = tc.repo.ValidateCredentials(&tourist)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "invalid credentials 2"} )
		log.Printf("error: %w", err)
		return
	}

	token, err := utils.GenerateToken(tourist.Email, tourist.ID)
	if err != nil {
		log.Printf("token generation failed for user %d: %v", tourist.ID, err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not authenticate user", "error": err} )
		return 
	}

	context.JSON(http.StatusOK, gin.H{"message": "login successfull", "token": token })
}

