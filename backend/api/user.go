package api

import (
	"app/pkg/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetUser godoc
// @Summary Получить пользователей по id
// @Description Метод позволяет получить всех пользователей
// @Success 200
// @Router /api/user/ [get]
func GetUser(ctx *gin.Context) {

	response, err := storage.DB.SelectUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Не удалось получить данные.", "err": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"result": response})
}
