package router

import (
	"context"
	"net/http"

	auth_m "gentools/genapi/internal/middleware/auth"
	logger_m "gentools/genapi/internal/middleware/logger"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
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

func NewRouterRun(url string, dbpool *pgxpool.Pool, ctx context.Context) {
	r := gin.New()

	r.Use(gin.Recovery())
	r.Use(logger_m.GenLogger())

	r.Use(auth_m.GenAuth(dbpool, ctx))

	r.GET("/users", get_users)
	r.Run(url)
}

func get_users(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}
