package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string
	Item      string
	Completed bool
}

var todos = []todo{
	{ID: "1", Item: "Clean Room", Completed: true},
	{ID: "2", Item: "Bath", Completed: false},
	{ID: "3", Item: "Breakfast", Completed: false},
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.Run("localhost:9090")
}
