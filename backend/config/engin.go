package config

import (
	"app/api"
	"app/pkg/storage"
	"app/pkg/storage/postgreStorage"
	"database/sql"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	"os"
	"time"
)

func InitEngine() *gin.Engine {
	//gin.SetMode(gin.ReleaseMode)
	db, err := sql.Open("postgres", fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=disable",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_DATABASE"),
		os.Getenv("DB_ADDRESS"),
	))
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	storage.DB = postgreStorage.NewPostgreStorage(db)

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		//AllowOrigins:     []string{"https://example.com", "http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	aRouter := r.Group("api")
	tRouter := aRouter.Group("/task")
	uRouter := aRouter.Group("/user")

	tRouter.POST("/sim/:count", api.SimActivity)
	tRouter.POST("/", api.CreateTask)
	tRouter.DELETE("/:id", api.DeleteTask)

	uRouter.GET("/", api.GetUser)

	return r
}
