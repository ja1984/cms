package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	firebase "firebase.google.com/go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/gobuffalo/packr"
	"github.com/ja1984/cogCMS/backend/config"
	"github.com/ja1984/cogCMS/backend/database"
	"github.com/ja1984/cogCMS/backend/middleware"
	"github.com/ja1984/cogCMS/backend/routes"
	"github.com/ja1984/cogCMS/backend/services"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
)

const (
	host     = "db"
	port     = 5432
	user     = "postgres"
	password = "cogCMS"
	dbname   = "postgres"
)

func main() {
	log.Print("Started cms backend 😎")

	if *config.Environment != "local" {
		setupFirebaseClient()
	}

	database.REDIS = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := database.REDIS.Ping().Err()

	if err != nil {
		log.Fatalf("Redis errored on ping err: %v", err)
	}

	log.Print("Redis started and pinged 🤩")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	database.SQL, err = sqlx.Connect("postgres", psqlInfo)

	defer database.SQL.Close()

	if err != nil {
		log.Fatalf("Database errored on connect err: %v", err)
	}

	log.Print("Database started and pinged 😎")

	runDatabaseMigrations()

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
		adminGroup := apiGroup.Group("admin", middleware.VerifyAuth())
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

func setupFirebaseClient() {
	var err error
	ctx := context.Background()
	app, err := firebase.NewApp(ctx, nil)

	if err != nil {
		log.Fatalf("error getting Auth client: %v", err)
	}

	//Temp stuff just to get it building , need to figure out of we should do with credentails for firebase auth
	// Think we will just accept all tokens :)
	client, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v", err)
	}

	services.FirebaseClient = client
}

func runDatabaseMigrations() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("failed to connect database for migration [%v]", err)
	}
	defer db.Close()

	migrations := &migrate.PackrMigrationSource{
		Box: packr.NewBox("./../../db/migrations"),
	}

	migrate.SetTable("migrations")

	n, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		log.Fatalf("failed running migration [%v]", err)
	}
	fmt.Printf("Applied %d migrations! 😇\n", n)
}
