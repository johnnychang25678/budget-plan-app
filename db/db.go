package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
)

func ConnectDB() *pgx.Conn {
	connString := fmt.Sprintf(
		"user=%s host=%s password=%s port=%s dbname=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		fmt.Println(err)
		log.Fatal("Unable to connect to datatbase")
	}
	fmt.Println("Connected to DB!")
	return conn
}
