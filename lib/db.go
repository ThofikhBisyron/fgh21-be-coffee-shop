package lib

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func DB() *pgx.Conn {

	conn, err := pgx.Connect(context.Background(), "postgresql://postgres:1@157.230.243.8:54321/konis_caffee?sslmode=disable")

	if err != nil {
		fmt.Println(err)
	}
	return conn
}
