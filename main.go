package main

import (
	"budget-plan-app/backend/db"
	"context"
	"fmt"
)

func main() {
	// server := gin.Default()
	// pgxConfig := pgx.ConnConfig{
	// 	Host:     "localhost",
	// 	Database: "budget_plan_app",
	// }

	conn := db.ConnectDB()

	defer conn.Close(context.Background())
	var id int
	var email string

	conn.QueryRow(context.Background(), "select id, email from member where id = $1", 1).Scan(&id, &email)
	fmt.Println(id, email)

	// server.GET("/", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, "hello")
	// })

	// routers.Routes(server)

	// server.Run(":8080")
}
