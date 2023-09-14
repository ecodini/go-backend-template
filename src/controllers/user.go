package controllers

import (
	"backend-template/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FindUsers(c *gin.Context) {
	var users []models.User

	models.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}
