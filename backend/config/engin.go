package config

import (
	"app/api"
	"app/pkg/storage"
	"app/pkg/storage/postgreStorage"
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
)

func InitEngine() *gin.Engine {
	db, err := sql.Open("postgres", "your_connection_string") // TODO добавить подключение к БД
	if err != nil {
		log.Fatal(err)
	}

	storage.DB = postgreStorage.NewPostgreStorage(db)

	r := gin.Default()

	aRouter := r.Group("api/")
	tRouter := aRouter.Group("task/")
	uRouter := aRouter.Group("user/")

	tRouter.POST("/", api.CreateTask)
	tRouter.DELETE("/:id", api.DeleteTask)

	uRouter.GET("/", api.GetUser)

	return r
}
