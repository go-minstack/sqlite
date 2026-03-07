package sqlite

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func NewDB(log *slog.Logger) (*gorm.DB, error) {
	dsn := os.Getenv("MINSTACK_DB_URL")
	if dsn == "" {
		return nil, fmt.Errorf("MINSTACK_DB_URL is not set")
	}

	return gorm.Open(sqlite.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 newGormLogger(log, 200*time.Millisecond),
	})
}
