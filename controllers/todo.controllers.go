package controllers

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stheven26/db"
	"github.com/stheven26/models"
)

type DataInput struct {
	GroupID uint   `json:"activity_group_id"`
	Title   string `json:"title"`
}

func CreateToDo(c *gin.Context) {
	var dataInput DataInput
	var dataActivities models.Activities
	connection := db.GetConnectionDB()

	if err := c.ShouldBindJSON(&dataInput); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Status":  "Failed",
			"Message": err.Error(),
		})
		return
	}

	if err := connection.Where(`id = ?`, dataInput.GroupID).First(&dataActivities).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
	} else {
		data := models.ToDo{
			GroupID:   dataInput.GroupID,
			Title:     dataInput.Title,
			IsActive:  true,
			Priority:  "very-high",
			CreatedAt: time.Now(),
			UpdateAt:  time.Now(),
			DeletedAt: sql.NullTime{},
		}

		connection.Create(&data)

		c.JSON(http.StatusOK, gin.H{
			"Status":  "Success",
			"Message": "Success",
			"data":    data,
		})
	}

}

func GetAllDataToDo(c *gin.Context) {
	var data []models.ToDo

	connection := db.GetConnectionDB()

	connection.Find(&data)

	c.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "Success",
		"data":    data,
	})
}

func GetDataToDoById(c *gin.Context) {
	var data models.ToDo
	connection := db.GetConnectionDB()

	if err := connection.Where(`id = ?`, c.Param("id")).First(&data).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "Success",
		"data":    data,
	})
}

func UpdateToDoById(c *gin.Context) {
	var dataInput map[string]string
	var data models.ToDo
	connection := db.GetConnectionDB()

	if err := c.ShouldBindJSON(&dataInput); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Failed",
			"message": err.Error(),
		})
		return
	}

	if err := connection.Where(`id = ?`, c.Param("id")).First(&data).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	dataUpdate := models.ToDo{
		Title: dataInput["title"],
	}

	connection.Model(&data).Updates(&dataUpdate)

	c.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "Success",
		"data":    data,
	})
}

func DeleteToDo(c *gin.Context) {
	var data models.ToDo

	connection := db.GetConnectionDB()

	if err := connection.Where(`id = ?`, c.Param("id")).First(&data).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
	}

	connection.Delete(&data)

	c.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "Success",
		"data":    DeletedStruct{},
	})
}
