package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	ginjwt "github.com/appleboy/gin-jwt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/gobuffalo/packr"
	"github.com/ja1984/cogCMS/backend/database"
	"github.com/ja1984/cogCMS/backend/routes"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
	jwt "gopkg.in/dgrijalva/jwt-go.v3"
)

const (
	host     = "db"
	port     = 5432
	user     = "postgres"
	password = "cogCMS"
	dbname   = "postgres"
)

type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// User demo
type User struct {
	UserName  string
	FirstName string
	LastName  string
}

type GoogleAuthRequst struct {
	Token string `json:"token"`
}

var identityKey = "id"

func main() {
	log.Print("Started cms backend 😎")

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

	authHandler := setupAuthMiddleware()
	// setupGoth()

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
		apiGroup.POST("admin/register", routes.RegisterUser)
		apiGroup.GET("auth/refresh_token", authHandler.RefreshHandler)
		apiGroup.POST("auth/login", authHandler.LoginHandler)

		apiGroup.POST("auth/google", func(c *gin.Context) {
			var reqest GoogleAuthRequst
			err := c.BindJSON(&reqest)

			if err != nil {
				log.Printf("Couldnt bind json auth/google, err: %v", err)
			}

			data := User{
				UserName:  "1",
				LastName:  "Pata",
				FirstName: "Data",
			}

			// Create the token
			token := jwt.New(jwt.GetSigningMethod(authHandler.SigningAlgorithm))
			claims := token.Claims.(jwt.MapClaims)

			if authHandler.PayloadFunc != nil {
				for key, value := range authHandler.PayloadFunc(data) {
					claims[key] = value
				}
			}

			expire := authHandler.TimeFunc().Add(authHandler.Timeout)
			claims["exp"] = expire.Unix()
			claims["orig_iat"] = authHandler.TimeFunc().Unix()
			tokenString, err := token.SignedString(authHandler.Key)

			if err != nil {
				c.Error(ginjwt.ErrFailedTokenCreation)
				c.Status(http.StatusUnauthorized)
				return
			}

			authHandler.LoginResponse(c, http.StatusOK, tokenString, time.Now())
		})

		r.NoRoute(authHandler.MiddlewareFunc(), func(c *gin.Context) {
			claims := ginjwt.ExtractClaims(c)
			log.Printf("NoRoute claims: %#v\n", claims)
			c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
		})

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

func setupAuthMiddleware() *ginjwt.GinJWTMiddleware {
	// the jwt middleware
	authMiddleware := &ginjwt.GinJWTMiddleware{
		Realm:      "test zone",
		Key:        []byte("secret key"),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		PayloadFunc: func(data interface{}) ginjwt.MapClaims {
			if v, ok := data.(*User); ok {
				return ginjwt.MapClaims{
					identityKey: v.UserName,
					"firstname": v.FirstName,
					"lastname":  v.LastName,
				}
			}
			return ginjwt.MapClaims{}
		},
		IdentityHandler: func(claims jwt.MapClaims) interface{} {
			return &User{
				UserName: claims["id"].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals login
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", ginjwt.ErrMissingLoginValues
			}
			userID := loginVals.Username
			password := loginVals.Password

			if (userID == "admin" && password == "admin") || (userID == "test" && password == "test") {
				return &User{
					UserName:  userID,
					LastName:  "Pata",
					FirstName: "Data",
				}, nil
			}

			return nil, ginjwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(*User); ok && v.UserName == "admin" {
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	}

	return authMiddleware
}

// func setupGoth() {
// 	goth.UseProviders(
// 		google.New(os.Getenv("GOOGLE_KEY"), os.Getenv("GOOGLE_SECRET"), "http://localhost:5000/auth/google/callback"),
// 	)
// }
