package internalsql

import (
	"database/sql"
	"fmt"
	"go-tweets/internal/config"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Connectmysql(cfg *config.Config) (*sql.DB, error) {
	databaseURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=%s",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName, "Asia%2FKolkata")
	db, err := sql.Open("mysql", databaseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MySQL: %w", err)
	}

	log.Println("Connected to MySQL database successfully")
	return db, nil
}
