package router

import (
	"net/http"

	"gentools/genapi/internal/middleware"

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

func NewRouterRun(url string) {
	r := gin.New()

	r.Use(gin.Recovery())
	r.Use(middleware.GenLogger())

	r.GET("/users", get_users)
	r.Run(url)
}

func get_users(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}
