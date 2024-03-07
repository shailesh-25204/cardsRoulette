package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// var (
// 	RedisAddr = "localhost:6379"
// )

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database, err := NewDatabase()
	if err != nil {
		log.Fatalf("Failed to connect to redis : %s", err.Error())
	}
	hub := newHub(database)
	go hub.run()
	router := initRouter(hub)
	// router := gin.Default()
	// router.GET("/", func(c *gin.Context) {
	// 	c.Header("Content-Type", "application/json")
	// 	serveWs(c.Writer, c.Request)
	// })

	router.Run(":8080")
}

func initRouter(hub *Hub) *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		serveWs(hub, c.Writer, c.Request)
	})
	return r
}
