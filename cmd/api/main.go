package main

import (
	// External

	"context"
	"os"

	// Internal
	"gentools/genapi/internal/core/db"
	"gentools/genapi/internal/core/logger"
	"gentools/genapi/internal/core/router"
)

func main() {
	// Initialize the logger
	logger.InitLogger()

	// Open a new global dbpool and close it when out of scope
	dbpool := db.OpenNewDbPool(os.Getenv("DB_URL"))
	defer dbpool.Close()

	ctx := context.Background()

	// Create and run the gin router with appropriate middlewares
	router.NewRouterRun("0.0.0.0:8080", dbpool, ctx)
}
