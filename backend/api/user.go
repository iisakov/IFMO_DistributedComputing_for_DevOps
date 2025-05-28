package api

import (
	"app/pkg/storage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetUser godoc
// @Summary Получить пользователей по id
// @Description Метод позволяет получить всех пользователей
// @Tags User
// @Param id path int true "Пользователь"
// @Accept json
// @Produce json
// @Success 200 {object} model.User{}
// @Router       /user/{id} [get]
func GetUser(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Что-то не так с параметром id.", "err": err.Error()})
	}

	response, err := storage.DB.SelectUser(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Не удалось получить данные.", "err": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"result": response})
}
