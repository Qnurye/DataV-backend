package db

import (
	"fmt"
	"log"

	"datav-api/config"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func InitDB() {
	cfg := config.AppConfig.DB
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database,
	)

	var err error
	DB, err = sqlx.Connect(cfg.Driver, dsn)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %s", err)
	}
	log.Println("Database connected successfully!")
}
