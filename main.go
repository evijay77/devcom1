package main

import (
	"context"
	"fmt"

	"github.com/vijaysoul/devcom/routes"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func main() {
	conn, err = connectDB()
	if err != nil {
		return
	}

	router := gin.Default()

	router.Use(dbMiddleware(*conn))

	usersGroup := router.Group("users")
	{
		usersGroup.POST("register", dbMiddleware, routes.UsersRegister)
	}

	router.Run(":3000")
}

func connectDB() (c *pgx.Conn, erro error) {
	conn, err := pgx.Connect(context.Background(), "postgresql://postgres:vijay@localhost:5432/devcom")
	if err != nil {
		fmt.Println(("Error connecting to DB"))
		fmt.Println(err.Error())
	}
	_ = conn.Ping(context.Background())
	return conn, err
}

func dbMiddleware(conn, pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", conn)
		c.Next()
	}
}
