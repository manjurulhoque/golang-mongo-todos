package controllers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/manjurulhoque/go-mongo-crud/database"
	"github.com/manjurulhoque/go-mongo-crud/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"time"
)

var todoCollection *mongo.Collection = database.OpenCollection(database.Client, "todos")
var validate = validator.New()

func GetTodos() gin.HandlerFunc {
	return func(c *gin.Context) {

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var todo models.Todo
		var todos []models.Todo

		cursor, err := todoCollection.Find(ctx, bson.D{})
		if err != nil {
			defer cursor.Close(ctx)
			c.JSON(http.StatusOK, todos)
		}

		for cursor.Next(ctx) {
			err := cursor.Decode(&todo)
			if err != nil {
				c.JSON(http.StatusOK, todos)
			}
			todos = append(todos, todo)
		}
		c.JSON(http.StatusOK, todos)
	}
}

func CreateTodo() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var todo models.Todo

		if err := c.BindJSON(&todo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			defer cancel()
			return
		}

		validationErr := validate.Struct(todo)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			defer cancel()
			return
		}
		defer cancel()
		todo.ID = primitive.NewObjectID()
		todo.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		todo.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

		result, insertErr := todoCollection.InsertOne(ctx, todo)
		if insertErr != nil {
			msg := fmt.Sprintf("Todo item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, result)
	}
}
