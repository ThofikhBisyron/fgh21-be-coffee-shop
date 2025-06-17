package lib

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func DB() *pgx.Conn {
	cfg := LoadConfig()
	conn, err := pgx.Connect(context.Background(), "postgresql://"+cfg.DBUser+":"+cfg.DBPassword+"@"+cfg.DBHost+":"+cfg.DBPort+"/"+cfg.DBName+"?sslmode=disable")

	if err != nil {
		fmt.Println(err)
	}
	return conn
}
