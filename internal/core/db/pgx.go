package db

import (
	"fmt"
)

func InitDB(db_url string) {
	fmt.Printf("Connecting to database at: %s\n", db_url)
}
