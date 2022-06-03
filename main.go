package main

import (
	"github.com/gin-gonic/gin"
	"github.com/manjurulhoque/go-mongo-crud/database"
	"github.com/manjurulhoque/go-mongo-crud/routes"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
)

var todoCollection *mongo.Collection = database.OpenCollection(database.Client, "todos")

func main() {
	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	router := gin.New()
	router.Use(gin.Logger())
	v1 := router.Group("/v1/api")
	routes.TodosRoutes(v1)
	//router.Use(middleware.Authentication())

	err := router.Run(":" + PORT)
	if err != nil {
		return
	}
}
