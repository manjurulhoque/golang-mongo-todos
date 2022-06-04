package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/manjurulhoque/go-mongo-crud/controllers"
)

func TodosRoutes(incomingRoutes *gin.RouterGroup) {
	incomingRoutes.GET("/todos", controllers.GetTodos())
	incomingRoutes.POST("/todos", controllers.CreateTodo())
}
