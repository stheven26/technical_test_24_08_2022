package routes

import (
	"technical_test_24_08_2022/controllers"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func routes() {
	router.POST("/activity-groups", controllers.CreateActivities)
	router.GET("/activity-groups", controllers.GetAllActivities)
	router.GET("/activity-groups/:id", controllers.GetActivitiesById)
	router.DELETE("/activity-groups/:id", controllers.DeleteActivities)

	router.POST("/todo-items", controllers.CreateToDo)
	router.GET("/todo-items", controllers.GetAllDataToDo)
	router.GET("/todo-items/:id", controllers.GetDataToDoById)
	router.PATCH("/todo-items/:id", controllers.UpdateToDoById)
	router.DELETE("/todo-items/:id", controllers.DeleteToDo)
}

func StartApplication() {
	routes()
	router.Run(":3030")
}
