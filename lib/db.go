package lib

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func DB() *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), "postgresql://postgres:1@postgres:54321/konis_caffee?sslmode=disable")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("DB connected successfully")
	return conn
}
