package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/ja1984/cogCMS/backend/database"
	"github.com/ja1984/cogCMS/backend/routes"
)

func main() {
	log.Print("Started cms backend 😎")

	database.REDIS = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := database.REDIS.Ping().Err()

	if err != nil {
		log.Fatalf("Redis errored on ping err: %v", err)
	}

	log.Print("Redis started and pinged 🤩")

	r := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true

	r.Use(cors.New(corsConfig))

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	apiGroup := r.Group("api/v1")
	{
		adminGroup := apiGroup.Group("admin")
		{
			adminGroup.GET("pages", routes.GetPages)
			adminGroup.POST("page", routes.CreatePage)

			adminGroup.POST("collection", routes.CreateCollection)
			adminGroup.GET("collections", routes.ListCollections)
			adminGroup.GET("collection/:key", routes.GetCollection)
			adminGroup.DELETE("collection/:key", routes.GetCollection)

		}
	}

	r.Run() // listen and serve on 0.0.0.0:8080
}
