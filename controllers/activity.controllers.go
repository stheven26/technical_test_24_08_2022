package controllers

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/stheven26/db"

	"github.com/gin-gonic/gin"
	"github.com/stheven26/models"
)

type DeletedStruct struct {
}

func CreateActivities(c *gin.Context) {
	var dataInput map[string]string

	if err := c.ShouldBindJSON(&dataInput); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Failed",
			"message": err.Error(),
			"data":    dataInput,
		})
		return
	}

	data := models.Activities{
		Email:     dataInput["email"],
		Title:     dataInput["title"],
		CreatedAt: time.Now(),
		UpdateAt:  time.Now(),
		DeletedAt: sql.NullTime{},
	}

	connection := db.GetConnectionDB()

	connection.Create(&data)

	c.JSON(http.StatusOK, gin.H{
		"Status":  "Success",
		"Message": "Success",
		"data":    data,
	})
}

func GetAllActivities(c *gin.Context) {
	var data []models.Activities

	connection := db.GetConnectionDB()

	connection.Find(&data)

	c.JSON(http.StatusOK, gin.H{
		"Status":  "Success",
		"Message": "Success",
		"data":    data,
	})
}

func GetActivitiesById(c *gin.Context) {
	var data models.Activities

	connection := db.GetConnectionDB()

	if err := connection.Where(`id = ?`, c.Param("id")).First(&data).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status":  "Failed",
			"Message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Status":  "Success",
		"Message": "Success",
		"data":    data,
	})
}

func DeleteActivities(c *gin.Context) {
	var data models.Activities

	id := c.Param("id")

	connection := db.GetConnectionDB()

	if err := connection.Where(`id = ?`, id).First(&data).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status":  "Failed",
			"Message": err.Error(),
		})
		return
	}

	connection.Delete(&data)

	c.JSON(http.StatusOK, gin.H{
		"Status":  "NotFound",
		"Message": "Activities With ID " + id + " Not Found",
		"data":    DeletedStruct{},
	})
}
