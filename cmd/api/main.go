package main

import (
	// External

	"os"

	// Internal
	"gentools/genapi/internal/core/db"
	"gentools/genapi/internal/core/logger"
	"gentools/genapi/internal/core/router"
)

func main() {
	logger.InitLogger()
	db.InitDB(os.Getenv("DB_URL"))

	router.NewRouterRun("0.0.0.0:8080")
}
