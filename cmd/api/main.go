package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var users = []user{
	{ID: "1", Name: "11"},
	{ID: "2", Name: "22"},
	{ID: "3", Name: "33"},
}

func main() {
	r := gin.Default()
	r.GET("/users", get_users)

	r.Run("0.0.0.0:8080")
}

func get_users(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}
