package sqlite

import (
	"fmt"
	"os"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func NewDB() (*gorm.DB, error) {
	dsn := os.Getenv("MINSTACK_DB_URL")
	if dsn == "" {
		return nil, fmt.Errorf("MINSTACK_DB_URL is not set")
	}

	return gorm.Open(sqlite.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})
}
